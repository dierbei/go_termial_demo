package cmd

import (
	"context"
	"log"
	"os"

	k8s "go_terminal/internal/pkg/kubernetes_client"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	podCmd.Flags().String("namespace", "default", "kubernetes namespace")
	RootCmd.AddCommand(podCmd)
}

var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "Print the pod name of namespace",
	RunE: func(cmd *cobra.Command, args []string) error {
		podList, err := k8s.GetK8sCli().CoreV1().Pods(cmd.Flag("namespace").Value.String()).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Println(err)
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"名称", "命名空间", "IP", "状态"})

		for _, pod := range podList.Items {
			table.Append([]string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase)})
		}

		table.Render()

		return nil
	},
}
