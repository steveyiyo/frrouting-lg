package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func lg(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// io.WriteString(w, "https://github.com/steveyiyo/frrouting-lg")

	r.ParseForm()
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		str := `<!DOCTYPE html>
		<html>
		
		<head>
		
			<title>Looking Glass</title>
		
			<!-- Metadata -->
			<meta charset='utf-8'>
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
		
			<!-- Facicon -->
			<link rel='shortcut icon' href='https://static.yiy.tw/media/logo/Yi_logo.png'>
		
			<!-- CSS -->
			<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/normalize.css@8.0.1/normalize.css"
				integrity="sha256-WAgYcAck1C1/zEl5sBl5cfyhxtLgKGdpI3oKyJffVRI=" crossorigin="anonymous">
		
			<!-- Style -->
			<style>
				* {
					font-family: "SF Pro TC", "SF Pro Text", "SF Pro Icons", "PingFang TC", "Helvetica Neue", "Helvetica", "Arial", "Microsoft JhengHei", wf_SegoeUI, "Segoe UI", Segoe, "Segoe WP", Tahoma, Verdana, Ubuntu, "Bitstream Vera Sans", "DejaVu Sans", Tahoma, 微軟正黑體, "LiHei Pro", "WenQuanYi Micro Hei", "Droid Sans Fallback", "AR PL UMing TW", Roboto, "Helvetica Neue", "Hiragino Maru Gothic ProN", メイリオ, "ヒラギノ丸ゴ ProN W4", Meiryo, "Droid Sans", sans-serif
				}
		
				html {
					height: 100%;
				}
		
				body {
					min-height: 100%;
					background-color: #F2F2F2;
				}
		
				body.dark {
					background-color: #222222 !important;
				}
		
				* {
					background: #F2F2F2;
					color: #222222;
				}
		
				.dark * {
					background: #222222 !important;
					color: #F2F2F2 !important;
				}
		
				body {
					display: flex;
					justify-content: center;
					align-items: center;
					flex-direction: column;
				}
		
				.container {
					transition: 0.6s;
					padding: 2em;
					border-radius: 5px;
					-webkit-box-shadow: 0px 2px 5px 0px rgba(0, 0, 0, 0.75);
					-moz-box-shadow: 0px 2px 5px 0px rgba(0, 0, 0, 0.75);
					box-shadow: 0px 2px 5px 0px rgba(0, 0, 0, 0.75);
					display: flex;
					flex-direction: column;
				}
		
				.container>* {
					margin-top: 0.5em;
		
					width: 100%;
				}
		
				button {
					padding: 1em;
					border-radius: 5px;
					border: 2px lightgray solid;
					transition: 0.4s;
					width: 100%;
					cursor: pointer;
					font-weight: bold;
				}
		
				button:hover {
					filter: invert(90%) hue-rotate(180deg);
					-webkit-box-shadow: 0px 1px 5px 0px rgba(0, 0, 0, 0.75);
					-moz-box-shadow: 0px 1px 5px 0px rgba(0, 0, 0, 0.75);
					box-shadow: 0px 1px 5px 0px rgba(0, 0, 0, 0.75);
				}
		
				input {}
		
				select {}
		
				[v-cloak] {
					opacity: 0;
				}
				
			</style>
		</head>
		
		<body>
			<h1> FRR Looking Glass </h1>
			<div class="container" id="app" v-cloak>
				<label>Server :</label>
				<select v-model="form.server">
					<option v-for="server in options.SERVERS" :key="server.value" :value="server.value">
						{{ server.name }}
					</option>
				</select>
				<label>Action :</label>
				<select v-model="form.action">
					<option v-for="action in options.ACTIONS" :key="action.value" :value="action.value">
						{{ action.name }}
					</option>
				</select>
				<label>Target :</label>
				<input v-model="form.target">
				<button @click="sendRequest">Start !</button>
			</form>
				<label>Result :</label>
				<pre>{{result}}</pre>
			</div>
		
			<!-- Scripts -->
			<script src="https://unpkg.com/vue@next"></script>
		
		
			<script>
				$('button').click(function (e) {
					e.preventDefault();
					var formData = {
						'IP': $('input[name=IP]').val(),
						'Action': $('select[name=Action]').val(),
					};
					$.ajax({
						url: '',
						data: formData,
						processData: true,
						type: "POST",
						encode: true,
						success: function (data) {
							console.log(data);
						}, error: function (data) {
							alert('An error occurred');
						}
					})
				});
			</script>
		
			<script>
				const ACTIONS = [
					{ name: "Ping", value: "ping" },
					{ name: "Traceroute", value: "traceroute" },
					{ name: "MTR", value: "mtr" },
					{ name: "BGP Summary", value: "bgpsummary" },
					{ name: "BGP Route V4", value: "routev4" },
					{ name: "BGP Route V6", value: "routev6" },
				]
		
				const SERVERS = [
					{ name: "JP", value: "jp-router" },
					{ name: "US", value: "us-router" },
					{ name: "EU", value: "eu-router" },
				]
		
				const SERVER_URL = "http://10.121.211.254:32280/"
		
				const app = {
					data() {
						return {
							options: {
								ACTIONS,
								SERVERS
							},
							form: {
								server: null,
								action: null,
								target: null
							},
							result: "",
							error: false
						}
					},
					methods: {
						sendRequest() {
							let formData = new FormData();
							formData.append('Action', this.form.action);
							formData.append('IP', this.form.target);
							fetch(SERVER_URL, {
								method: "POST",
								body: formData
							})
								.then(_ => _.text())
								.then(result => {
									this.result = result
								})
								.catch(error => {
									console.error(error)
									this.error = true
								})
						}
		
					}
				}
		
				Vue.createApp(app).mount('#app')
			</script>
		</body>
		
		</html>`
		w.Write([]byte(str))
	} else if r.Method == "POST" {
		r.ParseForm()
		Router := "NULL"
		Action := "NULL"
		IP := "NULL"

		Router = r.FormValue("Router")
		Action = r.FormValue("Action")
		IP = r.FormValue("IP")

		fmt.Println("Action: ", Action)
		fmt.Println("IP: ", IP)
		fmt.Println("Router: ", Router)

		resp, err := http.PostForm("http://10.121.211.254:32991",
			url.Values{"Action": {Action}, "IP": {IP}})
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(string(body))
		io.WriteString(w, string(body))
	}
}

func main() {
	http.HandleFunc("/", lg)

	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")
	fmt.Print("FRRouting Looking Glass\n")
	fmt.Print("Port listing at 32280\n")
	fmt.Print("Repo: https://github.com/steveyiyo/frrouting-lg\n")
	fmt.Print("Author: SteveYi\n")
	fmt.Print("Demo: https://network.steveyi.net/looking-glass\n")
	fmt.Print("\n")
	fmt.Print("-------------------\n")
	fmt.Print("\n")

	err := http.ListenAndServe(":32280", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
