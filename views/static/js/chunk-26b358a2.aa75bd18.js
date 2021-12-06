(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-26b358a2"],{"371d":function(t,e,o){"use strict";o.r(e);var l=function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("div",{staticClass:"app-container"},[o("div",{staticStyle:{"margin-bottom":"15px"}},[o("el-input",{attrs:{type:"textarea",rows:5,placeholder:t.$t("临时记录区域")},model:{value:t.formTmp.tempText,callback:function(e){t.$set(t.formTmp,"tempText",e)},expression:"formTmp.tempText"}})],1),o("el-form",{ref:"form",attrs:{model:t.form,"label-width":"120px"}},[o("el-form-item",{attrs:{label:t.$t("容器名称")}},[o("el-input",{attrs:{placeholder:t.$t("容器名称")},model:{value:t.form.Name,callback:function(e){t.$set(t.form,"Name",e)},expression:"form.Name"}})],1),o("el-form-item",{attrs:{label:t.$t("镜像")}},[o("el-input",{attrs:{placeholder:"docker.io/library/nginx:latest"},model:{value:t.form.Image,callback:function(e){t.$set(t.form,"Image",e)},expression:"form.Image"}})],1),o("el-form-item",{attrs:{label:t.$t("内存(M)")}},[o("el-input-number",{staticStyle:{width:"210px"},attrs:{controls:!1,placeholder:t.$t("内存，如：128")},model:{value:t.formTmp.Memory,callback:function(e){t.$set(t.formTmp,"Memory",e)},expression:"formTmp.Memory"}}),o("el-input",{staticStyle:{width:"50px"},attrs:{disabled:"",placeholder:"M"}})],1),o("el-form-item",{attrs:{label:t.$t("端口映射")}},[o("el-button",{attrs:{type:"success",size:"small",icon:"el-icon-plus"},on:{click:function(e){return t.addPort()}}},[t._v(t._s(t.$t("添加")))]),t._l(t.form.Ports,(function(e,l){return o("div",{key:e.key,staticStyle:{"margin-top":"10px"},attrs:{items:t.form.Ports}},[o("el-input",{staticStyle:{width:"200px"},attrs:{placeholder:"8080"},model:{value:e.PublicPort,callback:function(o){t.$set(e,"PublicPort",o)},expression:"item.PublicPort"}},[o("template",{slot:"prepend"},[t._v(t._s(e.IP)+":")])],2),o("span",{staticStyle:{color:"#2b2b2b"}},[t._v(" -> ")]),o("el-input",{staticStyle:{width:"200px"},attrs:{placeholder:"8080"},model:{value:e.PrivatePort,callback:function(o){t.$set(e,"PrivatePort",o)},expression:"item.PrivatePort"}},[o("el-select",{staticStyle:{width:"80px"},attrs:{slot:"append",placeholder:"TCP"},slot:"append",model:{value:e.Type,callback:function(o){t.$set(e,"Type",o)},expression:"item.Type"}},[o("el-option",{attrs:{label:"TCP",value:"tcp"}}),o("el-option",{attrs:{label:"UDP",value:"udp"}}),o("el-option",{attrs:{label:"SCTP",value:"sctp"}})],1)],1),o("el-button",{staticStyle:{"margin-left":"15px"},attrs:{type:"danger",size:"small",icon:"el-icon-delete",circle:""},on:{click:function(o){return t.delPort(e,l)}}})],1)}))],2),o("el-form-item",{attrs:{label:t.$t("卷映射")}},[o("el-button",{attrs:{type:"success",size:"small",icon:"el-icon-plus"},on:{click:function(e){return t.addVolume()}}},[t._v(t._s(t.$t("添加")))]),o("el-link",{staticStyle:{"margin-left":"10px"},attrs:{type:"info",underline:!1,href:"https://docs.docker.com/storage/bind-mounts/#start-a-container-with-a-bind-mount"}},[t._v(t._s(t.$t("Bind模式相当于docker -v参数。如何挂载卷？")))]),t._l(t.formTmp.Volumes,(function(e,l){return o("div",{key:e.hostPath,staticStyle:{"margin-top":"10px"},attrs:{items:t.formTmp.Volumes}},[o("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:t.$t("源地址(宿主机地址)")},model:{value:e.Source,callback:function(o){t.$set(e,"Source",o)},expression:"item.Source"}},[o("el-select",{staticStyle:{width:"100px"},attrs:{slot:"prepend",disabled:"",placeholder:t.$t("绑定方式")},slot:"prepend",model:{value:e.Type,callback:function(o){t.$set(e,"Type",o)},expression:"item.Type"}},[o("el-option",{attrs:{label:"Bind",value:"bind"}}),o("el-option",{attrs:{label:"Volume",value:"volume"}}),o("el-option",{attrs:{label:"Tmpfs",value:"tmpfs"}})],1)],1),o("span",{staticStyle:{color:"#2b2b2b"}},[t._v(" : ")]),o("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:t.$t("容器内地址")},model:{value:e.Destination,callback:function(o){t.$set(e,"Destination",o)},expression:"item.Destination"}},[o("el-select",{staticStyle:{width:"80px"},attrs:{slot:"append",clearable:"",placeholder:""},slot:"append",model:{value:e.RW,callback:function(o){t.$set(e,"RW",o)},expression:"item.RW"}},[o("el-option",{attrs:{label:t.$t("读写"),value:":rw"}}),o("el-option",{attrs:{label:t.$t("只读"),value:":ro"}})],1)],1),o("el-button",{staticStyle:{"margin-left":"15px"},attrs:{type:"danger",size:"small",icon:"el-icon-delete",circle:""},on:{click:function(o){return t.delVolume(e,l)}}})],1)}))],2),o("el-form-item",{attrs:{label:t.$t("环境变量")}},[o("el-button",{attrs:{type:"success",size:"small",icon:"el-icon-plus"},on:{click:function(e){return t.addEnv()}}},[t._v(t._s(t.$t("添加")))]),t._l(t.formTmp.Env,(function(e,l){return o("div",{key:l,staticStyle:{"margin-top":"10px"},attrs:{value:e,items:t.formTmp.Env}},[o("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:t.$t("变量key")},model:{value:e.key,callback:function(o){t.$set(e,"key",o)},expression:"item.key"}}),o("span",{staticStyle:{color:"#2b2b2b"}},[t._v(" = ")]),o("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:t.$t("变量值")},model:{value:e.value,callback:function(o){t.$set(e,"value",o)},expression:"item.value"}}),o("el-button",{staticStyle:{"margin-left":"15px"},attrs:{type:"danger",size:"small",icon:"el-icon-delete",circle:""},on:{click:function(o){return t.delEnv(e,l)}}})],1)}))],2),o("el-form-item",{attrs:{label:t.$t("日志驱动")}},[o("el-select",{attrs:{clearable:"",placeholder:t.$t("请选择")},model:{value:t.form.LogType,callback:function(e){t.$set(t.form,"LogType",e)},expression:"form.LogType"}},t._l(t.res.logTypes,(function(t){return o("el-option",{key:t.key,attrs:{label:t.name,value:t.name,title:t.summary}})})),1),o("el-link",{staticStyle:{"margin-left":"10px"},attrs:{type:"info",underline:!1,href:"https://docs.docker.com/config/containers/logging/configure/"}},[t._v(t._s(t.$t("如何选择日志驱动？")))])],1),o("el-form-item",{attrs:{label:t.$t("日志配置参数")}},[o("el-button",{attrs:{type:"success",size:"small",icon:"el-icon-plus"},on:{click:function(e){return t.addLogConfig()}}},[t._v(t._s(t.$t("添加")))]),t._l(t.formTmp.LogConfig,(function(e,l){return o("div",{key:e.key,staticStyle:{"margin-top":"10px"},attrs:{value:e,items:t.formTmp.LogConfig}},[o("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:t.$t("变量key")},model:{value:e.key,callback:function(o){t.$set(e,"key",o)},expression:"item.key"}}),o("span",{staticStyle:{color:"#2b2b2b"}},[t._v(" = ")]),o("el-input",{staticStyle:{width:"400px"},attrs:{placeholder:t.$t("变量值")},model:{value:e.value,callback:function(o){t.$set(e,"value",o)},expression:"item.value"}}),o("el-button",{staticStyle:{"margin-left":"15px"},attrs:{type:"danger",size:"small",icon:"el-icon-delete",circle:""},on:{click:function(o){return t.delLogConfig(e,l)}}})],1)}))],2),o("el-form-item",{attrs:{label:t.$t("目标服务器")}},[o("el-select",{staticClass:"filter-item",staticStyle:{width:"100%"},attrs:{multiple:"",filterable:"",placeholder:"ServerName",clearable:""},model:{value:t.form.ServerNames,callback:function(e){t.$set(t.form,"ServerNames",e)},expression:"form.ServerNames"}},t._l(t.res.ServerNames,(function(t){return o("el-option",{key:t,attrs:{label:t,value:t}})})),1)],1),o("el-form-item",[o("el-button",{staticStyle:{width:"200px"},attrs:{type:"primary"},on:{click:t.onSubmit}},[t._v(t._s(t.$t("发布")))])],1)],1)],1)},r=[],a=o("4dd0"),n={data:function(){return{res:{ServerNames:[],logTypes:[{name:"json-file",summary:"日志格式化为 JSON。这是 Docker 默认的日志驱动程序。"},{name:"syslog",summary:"将日志消息写入 syslog 工具。需配置syslog地址参数syslog-address，或守护程序必须在主机上运行。"},{name:"journald",summary:"将日志消息写入 journald。journald 守护程序必须在主机上运行。"},{name:"gelf",summary:"将日志消息写入 Graylog Extended Log Format (GELF) 终端，例如 Graylog 或 Logstash。"},{name:"fluentd",summary:"将日志消息写入 fluentd（forward input）。需配置fluentd地址参数fluentd-address，或fluentd守护程序必须在主机上运行。"},{name:"awslogs",summary:"将日志消息写入 Amazon CloudWatch Logs。"},{name:"splunk",summary:"Writes log messages to splunk using the HTTP Event Collector."},{name:"etwlogs",summary:"将日志消息写为 Windows 的 Event Tracing 事件。仅在Windows平台上可用。"},{name:"gcplogs",summary:"将日志消息写入 Google Cloud Platform (GCP) Logging。"},{name:"logentries",summary:"将日志消息写入 Rapid7 Logentries。"}],logTypes2:["json-file","fluentd","journald","syslog","gelf","awslogs","splunk","etwlogs","gcplogs","logentries"]},formTmp:{Ports:[],Env:[],Volumes:[],LogConfig:[],Memory:0,Replicas:1,tempText:""},form:{Image:"",Name:"",Ports:[],Volumes:[],Env:[],ServerNames:[],Memory:0,LogType:"",LogConfig:[]}}},created:function(){this.fetchData()},methods:{onSubmit:function(){var t=this;this.listLoading=!0;var e=[];for(var o in this.formTmp.Ports){var l=this.formTmp.Ports[o],r=l.IP;l.PublicPort&&l.PublicPort.length>0&&(r+=(r.length>0?":":"")+l.PublicPort),e.push((r.length>0?"->":"")+l.PrivatePort+"/"+l.Type)}var n=[];for(var s in this.formTmp.Env){var i=this.formTmp.Env[s];n.push(i.key+"="+i.value)}var m=[];for(var c in this.formTmp.Volumes){var u=this.formTmp.Volumes[c];m.push(u.Source+":"+u.Destination+u.RW)}var p={};for(var f in this.formTmp.LogConfig){var d=this.formTmp.LogConfig[f];p[d.key]=d.value}this.form.Env=n,this.form.Volumes=m,this.form.LogConfig=p,this.form.Memory=1024*this.formTmp.Memory*1024,console.log("this.formTmp:",this.formTmp),console.log("this.form:",this.form),Object(a["f"])(this.form).then((function(e){t.$message("发布成功"),t.listLoading=!1}))},onCancel:function(){this.$message({message:"cancel!",type:"warning"})},fetchData:function(){var t=this;this.listLoading=!0,Object(a["e"])().then((function(e){t.res.ServerNames=e.data,t.listLoading=!1}))},addPort:function(){this.form.Ports.push({PublicPort:"",PrivatePort:"",IP:"0.0.0.0",Type:"tcp"})},delPort:function(t,e){this.$delete(this.form.Ports,e)},addVolume:function(){this.formTmp.Volumes.push({Source:"",Destination:"",RW:"",Mode:"",Type:"bind",Propagation:"rprivate"})},delVolume:function(t,e){this.$delete(this.formTmp.Volumes,e)},addEnv:function(){this.formTmp.Env.push({key:"",value:""})},delEnv:function(t,e){this.$delete(this.formTmp.Env,e)},addLogConfig:function(){this.formTmp.LogConfig.push({key:"",value:""})},delLogConfig:function(t,e){this.$delete(this.formTmp.LogConfig,e)}}},s=n,i=(o("6d2f"),o("2877")),m=Object(i["a"])(s,l,r,!1,null,"539db848",null);e["default"]=m.exports},"4dd0":function(t,e,o){"use strict";o.d(e,"d",(function(){return r})),o.d(e,"e",(function(){return a})),o.d(e,"c",(function(){return n})),o.d(e,"b",(function(){return s})),o.d(e,"f",(function(){return i})),o.d(e,"a",(function(){return m}));var l=o("b775");function r(t){return Object(l["a"])({url:"/mgr/servers",method:"get",params:t})}function a(t){return Object(l["a"])({url:"/mgr/serverNames",method:"get",params:t})}function n(t){return Object(l["a"])({url:"/mgr/containers",method:"get",params:t})}function s(t){return Object(l["a"])({url:"/mgr/containerInfos",method:"get",params:t})}function i(t){return Object(l["a"])({url:"/mgr/publish",method:"post",data:t})}function m(t,e){return Object(l["a"])({url:"/mgr/container/"+t,method:"post",data:e})}},"6d2f":function(t,e,o){"use strict";o("6e82")},"6e82":function(t,e,o){}}]);