package main

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/rancher/wrangler/pkg/signals"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/harvester/harvester/pkg/cmd"
	harvesterconfig "github.com/harvester/harvester/pkg/config"
	apiserver "github.com/harvester/harvester/pkg/server"
	"github.com/harvester/harvester/pkg/webhook/config"
	"github.com/harvester/harvester/pkg/webhook/server"
)

func main() {
	var options config.Options

	flags := []cli.Flag{
		cli.IntFlag{
			Name:        "threadiness",
			EnvVar:      "THREADINESS",
			Usage:       "Specify controller threads",
			Value:       5,
			Destination: &options.Threadiness,
		},
		cli.IntFlag{
			Name:        "https-port",
			EnvVar:      "HARVESTER_WEBHOOK_SERVER_HTTPS_PORT",
			Usage:       "HTTPS listen port",
			Value:       9443,
			Destination: &options.HTTPSListenPort,
		},
		cli.StringFlag{
			Name:        "namespace",
			EnvVar:      "NAMESPACE",
			Destination: &options.Namespace,
			Usage:       "The harvester namespace",
			Required:    true,
		},
		cli.StringFlag{
			Name:        "controller-user",
			EnvVar:      "HARVESTER_CONTROLLER_USER_NAME",
			Destination: &options.HarvesterControllerUsername,
			Usage:       "The harvester controller username",
		},
		cli.StringFlag{
			Name:        "gc-user",
			EnvVar:      "GARBAGE_COLLECTION_USER_NAME",
			Destination: &options.GarbageCollectionUsername,
			Usage:       "The system username that performs garbage collection",
			Value:       "system:serviceaccount:kube-system:generic-garbage-collector",
		},
	}

	app := cmd.NewApp("Harvester Admission Webhook Server", "", flags, func(commonOptions *harvesterconfig.CommonOptions) error {
		return run(commonOptions, &options)
	})
	app.Run()
}

func run(commonOptions *harvesterconfig.CommonOptions, options *config.Options) error {
	logrus.Info("Starting webhook server")

	// Debug log messages for troubleshooting
	logrus.Debugf("Harvester: Starting run function")

	// Setup signal handling for graceful shutdown
	ctx := signals.SetupSignalContext()

	// More debug log messages
	logrus.Debugf("Harvester: Signal context setup complete")

	// Retrieve Kubernetes client configuration
	kubeConfig, err := apiserver.GetConfig(commonOptions.KubeConfig)
	if err != nil {
		return fmt.Errorf("failed to find kubeconfig: %v", err)
	}

	// More debug log messages
	logrus.Debugf("Harvester: Kubernetes configuration obtained from: %v", commonOptions.KubeConfig)

	// Retrieve Kubernetes REST client configuration
	restCfg, err := kubeConfig.ClientConfig()
	if err != nil {
		return fmt.Errorf("failed to get Kubernetes client config: %v", err)
	}

	// Create and start the webhook server
	s := server.New(ctx, restCfg, options)

	if err := s.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start webhook server: %v", err)
	}

	// Wait for signal to gracefully shutdown
	<-ctx.Done()

	// Debug log message to indicate graceful shutdown
	logrus.Debugf("Harvester: Shutdown completed")

	return nil
}
