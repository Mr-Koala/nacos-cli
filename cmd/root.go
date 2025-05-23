package cmd

import (
	"github/szpinc/nacosctl/pkg/nacos"
	"os"

	"github.com/spf13/cobra"
)

var namespace string
var group string
var username string // 新增：用于存储用户名
var password string // 新增：用于存储密码

var nacosClient *nacos.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nacosctl",
	Short: "nacos cli tools",
	Long:  `nacos cli tools`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	// PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

	// },
	ValidArgs: []string{"get", "delete", "edit"},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    // 从环境变量获取默认值
    defaultNamespace := os.Getenv("NACOS_NAMESPACE")
    defaultGroup := os.Getenv("NACOS_GROUP")
    if defaultGroup == "" {
        defaultGroup = "DEFAULT_GROUP"
    }

    // 设置命令行参数，使用环境变量作为默认值
    rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", defaultNamespace, "nacos namespace(or os env NACOS_NAMESPACE)")
    rootCmd.PersistentFlags().StringVarP(&group, "group", "g", defaultGroup, "nacos group(or os env NACOS_GROUP)")
    rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "nacos username(or os env NACOS_USER)")
    rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "nacos password(or os env NACOS_PASSWD)")

    // 只有当环境变量也没有设置时，才要求必填
    if defaultNamespace == "" {
        _ = rootCmd.MarkFlagRequired("namespace")
    }

    // 初始化客户端
    nacosClient = nacos.NewDefaultClient()
}
