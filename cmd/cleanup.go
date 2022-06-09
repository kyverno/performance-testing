/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"strconv"

	"github.com/husnialhamdani/kyvernop/objects"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup all resources created",
	Long:  `Cleanup all resources created`,
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetInt("size")
		cleanup(size, "default")
	},
}

func init() {
	cleanupCmd.Flags().IntP("size", "s", 0, "number of resources to cleanup")
	rootCmd.AddCommand(cleanupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func cleanup(size int, namespace string) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)
	log.Print("Cleaning up resources...")
	for i := size; i >= 0; i-- {
		counter := strconv.Itoa(i)
		objects.DeleteNamespace(*clientset, counter)
		objects.DeleteDeployment(*clientset, counter, namespace)
		objects.DeleteConfigmap(*clientset, counter, namespace)
		objects.DeletePod(*clientset, counter, namespace)
		objects.DeleteSecret(*clientset, counter, namespace)
		objects.DeleteCronjob(*clientset, counter, namespace)
	}
}
