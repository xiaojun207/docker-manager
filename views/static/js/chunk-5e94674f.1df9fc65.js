(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-5e94674f"],{"08a4":function(e,t,a){"use strict";a("33fa")},"1e63":function(e,t,a){"use strict";a.r(t);var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"app-container"},[a("el-tabs",{attrs:{type:"card"},model:{value:e.activeName,callback:function(t){e.activeName=t},expression:"activeName"}},[a("el-tab-pane",{attrs:{label:"表单模式",name:"2"}},[a("publish-form")],1),a("el-tab-pane",{attrs:{label:"Yaml模式",name:"3"}},[a("div",{staticStyle:{color:"#90949b","font-size":"14px",padding:"10px"}},[e._v("使用docker compose yaml格式规范")]),a("yaml-editor",{model:{value:e.yamlData,callback:function(t){e.yamlData=t},expression:"yamlData"}}),a("el-select",{staticClass:"filter-item",staticStyle:{"margin-top":"10px","margin-bottom":"10px","margin-right":"10px"},attrs:{multiple:"",filterable:"",placeholder:"ServerName",clearable:""},model:{value:e.form.ServerNames,callback:function(t){e.$set(e.form,"ServerNames",t)},expression:"form.ServerNames"}},e._l(e.res.ServerNames,(function(e){return a("el-option",{key:e,attrs:{label:e,value:e}})})),1),a("el-button",{staticStyle:{width:"200px"},attrs:{type:"primary"},on:{click:e.onSubmit}},[e._v(e._s(e.$t("发布")))])],1),a("el-tab-pane",{attrs:{label:"Json模式",name:"4",disabled:""}},[a("json-editor",{ref:"jsonEditor",model:{value:e.jsonData,callback:function(t){e.jsonData=t},expression:"jsonData"}})],1)],1)],1)},o=[],n=a("4dd0"),l=a("aa22"),s=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"yaml-editor"},[a("textarea",{ref:"textarea"})])},i=[],m=a("56b3"),c=a.n(m);a("0dd0"),a("a7be"),a("7a7a"),a("ced0"),a("8822"),a("b8d1");window.jsyaml=a("e2c1");var u={name:"YamlEditor",props:["value"],data:function(){return{yamlEditor:!1}},watch:{value:function(e){var t=this.yamlEditor.getValue();e!==t&&this.yamlEditor.setValue(this.value)}},mounted:function(){var e=this;this.yamlEditor=c.a.fromTextArea(this.$refs.textarea,{lineNumbers:!0,mode:"text/x-yaml",gutters:["CodeMirror-lint-markers"],theme:"monokai",lint:!0}),this.yamlEditor.setValue(this.value),this.yamlEditor.on("change",(function(t){e.$emit("changed",t.getValue()),e.$emit("input",t.getValue())}))},methods:{getValue:function(){return this.yamlEditor.getValue()}}},p=u,f=(a("08a4"),a("2877")),d=Object(f["a"])(p,s,i,!1,null,"6865aec9",null),y=d.exports,v=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"json-editor"},[a("textarea",{ref:"textarea"})])},h=[];a("acdf"),a("f9d4"),a("d2de");a("ae67");var b={name:"JsonEditor",props:["value"],data:function(){return{jsonEditor:!1}},watch:{value:function(e){var t=this.jsonEditor.getValue();e!==t&&this.jsonEditor.setValue(JSON.stringify(this.value,null,2))}},mounted:function(){var e=this;this.jsonEditor=c.a.fromTextArea(this.$refs.textarea,{lineNumbers:!0,mode:"application/json",gutters:["CodeMirror-lint-markers"],theme:"rubyblue",lint:!0}),this.jsonEditor.setValue(JSON.stringify(this.value,null,2)),this.jsonEditor.on("change",(function(t){e.$emit("changed",t.getValue()),e.$emit("input",t.getValue())}))},methods:{getValue:function(){return this.jsonEditor.getValue()}}},g=b,x=(a("1fd7"),Object(f["a"])(g,v,h,!1,null,"1958ddac",null)),k=x.exports,S=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"app-container"},[a("div",{staticStyle:{"margin-bottom":"15px"}},[a("el-input",{attrs:{type:"textarea",rows:5,placeholder:e.$t("临时记录区域")},model:{value:e.formTmp.tempText,callback:function(t){e.$set(e.formTmp,"tempText",t)},expression:"formTmp.tempText"}}),a("div",{staticStyle:{color:"#90949b","font-size":"14px",padding:"10px"}},[e._v("如下内容，可解析到表单："+e._s(e.tempPlaceholder))]),a("el-button",{staticStyle:{"margin-top":"10px"},attrs:{type:"primary",size:"mini"},on:{click:e.toForm}},[e._v(e._s(e.$t("解析")))])],1),a("el-form",{ref:"form",attrs:{model:e.form,"label-width":"120px"}},[a("el-form-item",{attrs:{label:e.$t("容器名称")}},[a("el-input",{attrs:{placeholder:e.$t("容器名称")},model:{value:e.form.Name,callback:function(t){e.$set(e.form,"Name",t)},expression:"form.Name"}}),a("el-checkbox",{model:{value:e.form.cover,callback:function(t){e.$set(e.form,"cover",t)},expression:"form.cover"}},[e._v(e._s(e.$t("如果容器名存在，则覆盖")))])],1),a("el-form-item",{attrs:{label:e.$t("镜像")}},[a("el-input",{attrs:{placeholder:"docker.io/library/nginx:latest"},model:{value:e.form.Image,callback:function(t){e.$set(e.form,"Image",t)},expression:"form.Image"}})],1),a("el-form-item",{attrs:{label:e.$t("内存(M)")}},[a("el-input-number",{staticStyle:{width:"210px"},attrs:{controls:!1,placeholder:e.$t("内存，如：128")},model:{value:e.formTmp.Memory,callback:function(t){e.$set(e.formTmp,"Memory",t)},expression:"formTmp.Memory"}}),a("el-input",{staticStyle:{width:"50px"},attrs:{disabled:"",placeholder:"M"}})],1),a("el-form-item",{attrs:{label:e.$t("端口映射")}},[a("el-button",{attrs:{type:"success",size:"small",icon:"el-icon-plus"},on:{click:function(t){return e.addPort()}}},[e._v(e._s(e.$t("添加")))]),e._l(e.form.Ports,(function(t,r){return a("div",{key:t.key,staticStyle:{"margin-top":"10px"},attrs:{items:e.form.Ports}},[a("el-input",{staticStyle:{width:"200px"},attrs:{placeholder:"8080"},model:{value:t.PublicPort,callback:function(a){e.$set(t,"PublicPort",a)},expression:"item.PublicPort"}},[a("template",{slot:"prepend"},[e._v(e._s(t.IP)+":")])],2),a("span",{staticStyle:{color:"#2b2b2b"}},[e._v(" -> ")]),a("el-input",{staticStyle:{width:"200px"},attrs:{placeholder:"8080"},model:{value:t.PrivatePort,callback:function(a){e.$set(t,"PrivatePort",a)},expression:"item.PrivatePort"}},[a("el-select",{staticStyle:{width:"80px"},attrs:{slot:"append",placeholder:"TCP"},slot:"append",model:{value:t.Type,callback:function(a){e.$set(t,"Type",a)},expression:"item.Type"}},[a("el-option",{attrs:{label:"TCP",value:"tcp"}}),a("el-option",{attrs:{label:"UDP",value:"udp"}}),a("el-option",{attrs:{label:"SCTP",value:"sctp"}})],1)],1),a("el-button",{staticStyle:{"margin-left":"15px"},attrs:{type:"danger",size:"small",icon:"el-icon-delete",circle:""},on:{click:function(a){return e.delPort(t,r)}}})],1)}))],2),a("el-form-item",{attrs:{label:e.$t("卷映射")}},[a("el-button",{attrs:{type:"success",size:"small",icon:"el-icon-plus"},on:{click:function(t){return e.addVolume()}}},[e._v(e._s(e.$t("添加")))]),a("el-link",{staticStyle:{"margin-left":"10px"},attrs:{type:"info",underline:!1,href:"https://docs.docker.com/storage/bind-mounts/#start-a-container-with-a-bind-mount"}},[e._v(" "+e._s(e.$t("Bind模式相当于docker -v参数。如何挂载卷？"))+" ")]),e._l(e.formTmp.Volumes,(function(t,r){return a("div",{key:t.hostPath,staticStyle:{"margin-top":"10px"},attrs:{items:e.formTmp.Volumes}},[a("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:e.$t("源地址(宿主机地址)")},model:{value:t.Source,callback:function(a){e.$set(t,"Source",a)},expression:"item.Source"}},[a("el-select",{staticStyle:{width:"100px"},attrs:{slot:"prepend",disabled:"",placeholder:e.$t("绑定方式")},slot:"prepend",model:{value:t.Type,callback:function(a){e.$set(t,"Type",a)},expression:"item.Type"}},[a("el-option",{attrs:{label:"Bind",value:"bind"}}),a("el-option",{attrs:{label:"Volume",value:"volume"}}),a("el-option",{attrs:{label:"Tmpfs",value:"tmpfs"}})],1)],1),a("span",{staticStyle:{color:"#2b2b2b"}},[e._v(" : ")]),a("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:e.$t("容器内地址")},model:{value:t.Destination,callback:function(a){e.$set(t,"Destination",a)},expression:"item.Destination"}},[a("el-select",{staticStyle:{width:"80px"},attrs:{slot:"append",clearable:"",placeholder:""},slot:"append",model:{value:t.RW,callback:function(a){e.$set(t,"RW",a)},expression:"item.RW"}},[a("el-option",{attrs:{label:e.$t("读写"),value:":rw"}}),a("el-option",{attrs:{label:e.$t("只读"),value:":ro"}})],1)],1),a("el-button",{staticStyle:{"margin-left":"15px"},attrs:{type:"danger",size:"small",icon:"el-icon-delete",circle:""},on:{click:function(a){return e.delVolume(t,r)}}})],1)}))],2),a("el-form-item",{attrs:{label:e.$t("环境变量")}},[a("el-button",{attrs:{type:"success",size:"small",icon:"el-icon-plus"},on:{click:function(t){return e.addEnv()}}},[e._v(e._s(e.$t("添加")))]),e._l(e.formTmp.Env,(function(t,r){return a("div",{key:r,staticStyle:{"margin-top":"10px"},attrs:{value:t,items:e.formTmp.Env}},[a("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:e.$t("变量key")},model:{value:t.key,callback:function(a){e.$set(t,"key",a)},expression:"item.key"}}),a("span",{staticStyle:{color:"#2b2b2b"}},[e._v(" = ")]),a("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:e.$t("变量值")},model:{value:t.value,callback:function(a){e.$set(t,"value",a)},expression:"item.value"}}),a("el-button",{staticStyle:{"margin-left":"15px"},attrs:{type:"danger",size:"small",icon:"el-icon-delete",circle:""},on:{click:function(a){return e.delEnv(t,r)}}})],1)}))],2),a("el-form-item",{attrs:{label:e.$t("日志驱动")}},[a("el-select",{attrs:{clearable:"",placeholder:e.$t("请选择")},model:{value:e.form.LogType,callback:function(t){e.$set(e.form,"LogType",t)},expression:"form.LogType"}},e._l(e.res.logTypes,(function(e){return a("el-option",{key:e.key,attrs:{label:e.name,value:e.name,title:e.summary}})})),1),a("el-link",{staticStyle:{"margin-left":"10px"},attrs:{type:"info",underline:!1,href:"https://docs.docker.com/config/containers/logging/configure/"}},[e._v(" "+e._s(e.$t("如何选择日志驱动？"))+" ")])],1),a("el-form-item",{attrs:{label:e.$t("日志配置参数")}},[a("el-button",{attrs:{type:"success",size:"small",icon:"el-icon-plus"},on:{click:function(t){return e.addLogConfig()}}},[e._v(e._s(e.$t("添加")))]),e._l(e.formTmp.LogConfig,(function(t,r){return a("div",{key:t.key,staticStyle:{"margin-top":"10px"},attrs:{value:t,items:e.formTmp.LogConfig}},[a("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:e.$t("变量key")},model:{value:t.key,callback:function(a){e.$set(t,"key",a)},expression:"item.key"}}),a("span",{staticStyle:{color:"#2b2b2b"}},[e._v(" = ")]),a("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:e.$t("变量值")},model:{value:t.value,callback:function(a){e.$set(t,"value",a)},expression:"item.value"}}),a("el-button",{staticStyle:{"margin-left":"15px"},attrs:{type:"danger",size:"small",icon:"el-icon-delete",circle:""},on:{click:function(a){return e.delLogConfig(t,r)}}})],1)}))],2),a("el-form-item",{attrs:{label:e.$t("目标主机")}},[a("el-select",{staticClass:"filter-item",staticStyle:{width:"100%"},attrs:{multiple:"",filterable:"",placeholder:"ServerName",clearable:""},model:{value:e.form.ServerNames,callback:function(t){e.$set(e.form,"ServerNames",t)},expression:"form.ServerNames"}},e._l(e.res.ServerNames,(function(e){return a("el-option",{key:e,attrs:{label:e,value:e}})})),1)],1),a("el-form-item",[a("el-button",{staticStyle:{width:"200px"},attrs:{type:"primary"},on:{click:e.onSubmit}},[e._v(e._s(e.$t("发布")))])],1)],1)],1)},T=[],$=(a("ac1f"),a("1276"),a("caad"),a("2ca0"),a("8a79"),{name:"PublishForm",data:function(){return{res:{ServerNames:[],logTypes:[{name:"json-file",summary:"日志格式化为 JSON。这是 Docker 默认的日志驱动程序。"},{name:"syslog",summary:"将日志消息写入 syslog 工具。需配置syslog地址参数syslog-address，或守护程序必须在主机上运行。"},{name:"journald",summary:"将日志消息写入 journald。journald 守护程序必须在主机上运行。"},{name:"gelf",summary:"将日志消息写入 Graylog Extended Log Format (GELF) 终端，例如 Graylog 或 Logstash。"},{name:"fluentd",summary:"将日志消息写入 fluentd（forward input）。需配置fluentd地址参数fluentd-address，或fluentd守护程序必须在主机上运行。"},{name:"awslogs",summary:"将日志消息写入 Amazon CloudWatch Logs。"},{name:"splunk",summary:"Writes log messages to splunk using the HTTP Event Collector."},{name:"etwlogs",summary:"将日志消息写为 Windows 的 Event Tracing 事件。仅在Windows平台上可用。"},{name:"gcplogs",summary:"将日志消息写入 Google Cloud Platform (GCP) Logging。"},{name:"logentries",summary:"将日志消息写入 Rapid7 Logentries。"}],logTypes2:["json-file","fluentd","journald","syslog","gelf","awslogs","splunk","etwlogs","gcplogs","logentries"]},tempPlaceholder:"docker run -d --name test-openresty -p 80:80 -p 443:443 -m 512m -e user=root -v /tmp/openresty/nginx/logs:/usr/local/openresty/nginx/logs openresty/openresty:latest",formTmp:{Ports:[],Env:[],Volumes:[],LogConfig:[],Memory:0,Replicas:1,tempText:""},form:{Image:"",Cmd:"",Name:"",Ports:[],Volumes:[],Env:[],ServerNames:[],Memory:0,LogType:"",LogConfig:[],cover:!0}}},created:function(){this.fetchData()},methods:{onSubmit:function(){var e=this;this.listLoading=!0;var t=[];for(var a in this.formTmp.Ports){var r=this.formTmp.Ports[a],o=r.IP;r.PublicPort&&r.PublicPort.length>0&&(o+=(o.length>0?":":"")+r.PublicPort),t.push((o.length>0?"->":"")+r.PrivatePort+"/"+r.Type)}var l=[];for(var s in this.formTmp.Env){var i=this.formTmp.Env[s];l.push(i.key+"="+i.value)}var m=[];for(var c in this.formTmp.Volumes){var u=this.formTmp.Volumes[c];m.push(u.Source+":"+u.Destination+u.RW)}var p={};for(var f in this.formTmp.LogConfig){var d=this.formTmp.LogConfig[f];p[d.key]=d.value}this.form.Env=l,this.form.Volumes=m,this.form.LogConfig=p,this.form.Memory=1024*this.formTmp.Memory*1024,console.log("this.form:",this.form),Object(n["e"])(this.form).then((function(t){e.$message("发布成功"),e.listLoading=!1}))},onCancel:function(){this.$message({message:"cancel!",type:"warning"})},fetchData:function(){var e=this;this.listLoading=!0,Object(l["d"])().then((function(t){e.res.ServerNames=t.data,e.listLoading=!1}))},addPort:function(){this.form.Ports.push({PublicPort:"",PrivatePort:"",IP:"0.0.0.0",Type:"tcp"})},delPort:function(e,t){this.$delete(this.form.Ports,t)},addVolume:function(){this.formTmp.Volumes.push({Source:"",Destination:"",RW:":rw",Mode:"",Type:"bind",Propagation:"rprivate"})},delVolume:function(e,t){this.$delete(this.formTmp.Volumes,t)},addEnv:function(){this.formTmp.Env.push({key:"",value:""})},delEnv:function(e,t){this.$delete(this.formTmp.Env,t)},addLogConfig:function(){this.formTmp.LogConfig.push({key:"",value:""})},delLogConfig:function(e,t){this.$delete(this.formTmp.LogConfig,t)},clear:function(){this.form.Image="",this.form.Ports=[],this.formTmp.Volumes=[],this.formTmp.Env=[]},toForm:function(){var e=this.formTmp.tempText,t=["run","-d","--name","-p","-e","-v","-m"],a=e.split(" "),r="",o="",n=!1,l=!0,s="",i="";for(var m in this.clear(),a)if(l||t.includes(a[m])){if(l=t.includes(a[m]),o+=a[m],""===a[m]){if(!n)continue;o+=" "}if("--name"===r)this.form.Name=o;else if("-m"===r){var c=o.toLowerCase().replaceAll("m","");this.formTmp.Memory=parseInt(c)}else if("-p"===r){var u=o.split(":");this.form.Ports.push({PublicPort:u[0],PrivatePort:u[1],IP:"0.0.0.0",Type:"tcp"})}else if("-v"===r){var p=o.split(":");this.formTmp.Volumes.push({Source:p[0],Destination:p[1],RW:":rw",Mode:"",Type:"bind",Propagation:"rprivate"})}else if("-e"===r){var f=o.split("=");this.formTmp.Env.push({key:f[0],value:f[1]})}(o.startsWith('"')||o.endsWith('"'))&&(n=!n),n||(r=o,o="")}else""===s?(s=a[m],this.form.Image=s):(i+=" "+a[m],this.form.Cmd=i)}}}),_=$,P=(a("a584"),Object(f["a"])(_,S,T,!1,null,"8aefc30e",null)),E=P.exports,w='version: "2"\n\nservices:\n  openresty:\n    container_name: test-openresty\n    image: openresty/openresty:latest\n    restart: always\n    ports:\n      - "80:80"\n      - "443:443"\n    environment:\n      - GET_HOSTS_FROM=dns\n    volumes:\n      - "/tmp/data/web:/data/web"\n      - "/tmp/openresty/nginx/logs:/usr/local/openresty/nginx/logs"\n    \n\n',j='[{"items":[{"market_type":"forexdata","symbol":"XAUUSD"},{"market_type":"forexdata","symbol":"UKOIL"},{"market_type":"forexdata","symbol":"CORN"}],"name":""},{"items":[{"market_type":"forexdata","symbol":"XAUUSD"},{"market_type":"forexdata","symbol":"XAGUSD"},{"market_type":"forexdata","symbol":"AUTD"},{"market_type":"forexdata","symbol":"AGTD"}],"name":"贵金属"},{"items":[{"market_type":"forexdata","symbol":"CORN"},{"market_type":"forexdata","symbol":"WHEAT"},{"market_type":"forexdata","symbol":"SOYBEAN"},{"market_type":"forexdata","symbol":"SUGAR"}],"name":"农产品"},{"items":[{"market_type":"forexdata","symbol":"UKOIL"},{"market_type":"forexdata","symbol":"USOIL"},{"market_type":"forexdata","symbol":"NGAS"}],"name":"能源化工"}]',N={components:{YamlEditor:y,JsonEditor:k,PublishForm:E},data:function(){return{res:{ServerNames:[]},form:{ServerNames:[]},activeName:"2",yamlData:w,jsonData:JSON.parse(j)}},created:function(){this.fetchData()},methods:{onSubmit:function(){var e=this,t={serverNames:this.form.ServerNames,yaml:this.yamlData};Object(n["f"])(t).then((function(t){e.$message("发布成功")}))},fetchData:function(){var e=this;this.listLoading=!0,Object(l["d"])().then((function(t){e.res.ServerNames=t.data,e.listLoading=!1}))}}},C=N,L=(a("2961"),Object(f["a"])(C,r,o,!1,null,"0a0f9b5d",null));t["default"]=L.exports},"1fd7":function(e,t,a){"use strict";a("877b")},2961:function(e,t,a){"use strict";a("f905")},"33fa":function(e,t,a){},"4dd0":function(e,t,a){"use strict";a.d(t,"d",(function(){return o})),a.d(t,"b",(function(){return n})),a.d(t,"c",(function(){return l})),a.d(t,"e",(function(){return s})),a.d(t,"f",(function(){return i})),a.d(t,"a",(function(){return m}));var r=a("b775");function o(e){return Object(r["a"])({url:"/mgr/container/list",method:"get",params:e})}function n(e){return Object(r["a"])({url:"/mgr/container/detail",method:"get",params:e})}function l(e){return Object(r["a"])({url:"/mgr/containerInfos",method:"get",params:e})}function s(e){return Object(r["a"])({url:"/mgr/publish",method:"post",data:e})}function i(e){return Object(r["a"])({url:"/mgr/publish/yaml",method:"post",data:e})}function m(e,t){return Object(r["a"])({url:"/mgr/container/"+e,method:"post",data:t})}},"7d43":function(e,t,a){},"877b":function(e,t,a){},a584:function(e,t,a){"use strict";a("7d43")},aa22:function(e,t,a){"use strict";a.d(t,"c",(function(){return o})),a.d(t,"d",(function(){return n})),a.d(t,"b",(function(){return l})),a.d(t,"a",(function(){return s}));var r=a("b775");function o(e){return Object(r["a"])({url:"/mgr/server/list",method:"get",params:e})}function n(e){return Object(r["a"])({url:"/mgr/server/names",method:"get",params:e})}function l(e){return Object(r["a"])({url:"/mgr/server/detail",method:"get",params:e})}function s(e){return Object(r["a"])({url:"/mgr/server/delete",method:"post",data:e})}},f905:function(e,t,a){}}]);