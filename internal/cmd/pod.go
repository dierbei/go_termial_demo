package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	k8s "go_terminal/internal/pkg/kubernetes_client"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	podLabels bool
)

func init() {

	podCmd.Flags().String("namespace", "default", "kubernetes namespace")
	podCmd.Flags().BoolVarP(&podLabels, "labels", "", false, "kubernetes namespace")
	RootCmd.AddCommand(podCmd)
}

var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "Print the pod name of namespace",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println(cmd.Flags().GetString("namespace"))
		//log.Println(cmd.Flag("namespace").Value.String())
		// cmd.Flag("namespace").Value.String()
		podList, err := k8s.GetK8sCli().CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Println(err)
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"名称", "命名空间", "IP", "状态"})

		for _, pod := range podList.Items {
			if podLabels {
				table.Append([]string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase), Map2String(pod.Labels)})
			} else {
				table.Append([]string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase)})
			}

		}

		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(true)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetRowSeparator("")

		table.Render()

		return nil
	},
}

func Map2String(m map[string]string) (ret string) {
	for k, v := range m {
		ret += fmt.Sprintf("%s=%s\n", k, v)
	}
	return
}
