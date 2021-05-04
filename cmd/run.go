/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xiaomi388/virtual-music-system/metadata"
	"github.com/xiaomi388/virtual-music-system/metadata/driver/netease"
	metaIntf "github.com/xiaomi388/virtual-music-system/metadata/intf"
	"github.com/xiaomi388/virtual-music-system/song"
	songDriver "github.com/xiaomi388/virtual-music-system/song/driver"
	songIntf "github.com/xiaomi388/virtual-music-system/song/intf"
	"github.com/xiaomi388/virtual-music-system/song/model"
	"net/http"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ge := gin.Default()
		songHS := songIntf.HTTPService{
			Service: &song.Service{
				LocalRepo: &songDriver.FileRepository{
					RootDir: viper.GetString("song.local-repo.file.rootdir"),
				},
				RemoteRepos: []model.RemoteRepository{
					&songDriver.YoutubeRepository{},
				},
			},
			GE: ge,
		}
		metaHS := metaIntf.HTTPService{
			Service: &metadata.Service{
				SongRepo: &netease.SongRepository{
					BaseURL: viper.GetString("metadata.driver.netease.baseurl"),
					Client:  http.Client{},
				},
			},
			GE: ge,
		}
		songHS.Register()
		metaHS.Register()
		ge.Run(viper.GetString("http.hostport"))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
