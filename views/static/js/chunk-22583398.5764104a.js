(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-22583398"],{"4dd0":function(t,e,n){"use strict";n.d(e,"d",(function(){return a})),n.d(e,"b",(function(){return i})),n.d(e,"c",(function(){return o})),n.d(e,"e",(function(){return s})),n.d(e,"a",(function(){return l}));var r=n("b775");function a(t){return Object(r["a"])({url:"/mgr/container/list",method:"get",params:t})}function i(t){return Object(r["a"])({url:"/mgr/container/detail",method:"get",params:t})}function o(t){return Object(r["a"])({url:"/mgr/containerInfos",method:"get",params:t})}function s(t){return Object(r["a"])({url:"/mgr/publish",method:"post",data:t})}function l(t,e){return Object(r["a"])({url:"/mgr/container/"+t,method:"post",data:e})}},a352:function(t,e,n){"use strict";n.r(e);var r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"app-container"},[n("div",{staticClass:"filter-container"},[n("el-select",{staticClass:"filter-item",staticStyle:{width:"300px"},attrs:{multiple:"",filterable:"",placeholder:t.$t("服务器"),clearable:"","collapse-tags":""},model:{value:t.listQuery.ServerNames,callback:function(e){t.$set(t.listQuery,"ServerNames",e)},expression:"listQuery.ServerNames"}},t._l(t.res.serverNames,(function(t){return n("el-option",{key:t,attrs:{label:t,value:t}})})),1),n("el-select",{staticClass:"filter-item",staticStyle:{width:"300px","margin-left":"10px"},attrs:{multiple:"",filterable:"",placeholder:t.$t("容器名称"),clearable:"","collapse-tags":""},model:{value:t.listQuery.ContainerNames,callback:function(e){t.$set(t.listQuery,"ContainerNames",e)},expression:"listQuery.ContainerNames"}},t._l(t.res.ContainerNames,(function(t){return n("el-option",{key:t,attrs:{label:t,value:t}})})),1),n("el-select",{staticClass:"filter-item",staticStyle:{width:"140px","margin-left":"10px","margin-right":"10px"},attrs:{placeholder:"state",clearable:"",filterable:""},model:{value:t.listQuery.state,callback:function(e){t.$set(t.listQuery,"state",e)},expression:"listQuery.state"}},[n("el-option",{key:"running",attrs:{label:"running",value:"running"}}),n("el-option",{key:"exited",attrs:{label:"exited",value:"exited"}})],1),n("el-button",{staticClass:"filter-item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:t.fetchData}},[t._v(" Search ")])],1),n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.listLoading,expression:"listLoading"}],staticStyle:{"margin-top":"20px"},attrs:{data:t.list,"span-method":t.spanMethod,"element-loading-text":"Loading",border:"",fit:"","highlight-current-row":""}},[n("el-table-column",{attrs:{align:"center",label:"ID",width:"65"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.$index+1)+" ")]}}])}),n("el-table-column",{attrs:{label:"ServerName",width:"170"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.ServerName)+" ")]}}])}),n("el-table-column",{attrs:{label:"Name"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-button",{attrs:{type:"text",title:e.row.Name},on:{click:function(n){return t.openDetail(e.row)}}},[t._v(t._s(e.row.Name))])]}}])}),n("el-table-column",{attrs:{label:"IMAGE",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("span",[t._v(t._s(e.row.Image))])]}}])}),n("el-table-column",{attrs:{label:"Status",width:"200",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.Status)+" ")]}}])}),n("el-table-column",{attrs:{label:"State",width:"80",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return["running"===e.row.State?n("span",{staticStyle:{color:"#03c961"}},[t._v(t._s(e.row.State))]):t._e(),"running"!==e.row.State?n("span",{staticStyle:{color:"#d70404"}},[t._v(t._s(e.row.State))]):t._e()]}}])}),n("el-table-column",{attrs:{label:"Ports",width:"270",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return t._l(e.row.Ports,(function(e){return n("div",{key:e.key},[t._v(t._s(t.PortToStr(e)))])}))}}])}),n("el-table-column",{attrs:{align:"center",prop:"created_at",label:"Created",width:"200"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("i",{staticClass:"el-icon-time"}),n("span",[t._v(t._s(t.formatDate(e.row.Created)))])]}}])}),n("el-table-column",{attrs:{"class-name":"status-col",label:t.$t("操作"),width:"170",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-button",{attrs:{loading:t.listLoading,disabled:"running"!==e.row.State,type:"text"},on:{click:function(n){return t.ContainerOperator("stop",e.row)}}},[t._v(t._s(t.$t("停止")))]),n("el-button",{attrs:{loading:t.listLoading,type:"text"},on:{click:function(n){return t.ContainerOperator("remove",e.row)}}},[t._v(t._s(t.$t("删除")))]),n("el-button",{attrs:{loading:t.listLoading,type:"text"},on:{click:function(n){return t.ContainerOperator("restart",e.row)}}},[t._v(t._s(t.$t("重启")))])]}}])})],1),n("el-dialog",{attrs:{visible:t.dialogDetailVisible,title:t.$t("详情")},on:{"update:visible":function(e){t.dialogDetailVisible=e}}},[n("pre",[t._v(t._s(JSON.stringify(t.selectRow,null,2))+"\n      ")])])],1)},a=[],i=(n("d3b7"),n("6062"),n("3ca3"),n("ddb0"),n("d81d"),n("4dd0")),o=n("aa22"),s=n("ed08"),l=n("b1cf"),u={filters:{statusFilter:function(t){var e={published:"success",draft:"gray",deleted:"danger"};return e[t]}},data:function(){return{list:[],groupList:[],groups:{},listLoading:!1,dialogDetailVisible:!1,selectRow:{},res:{serverNames:[],ContainerNames:[],containerInfos:[]},listQuery:{ServerNames:[],ContainerNames:[],state:void 0}}},computed:{},created:function(){this.fetchData(),this.fetchContainerInfoData(),this.fetchServerNames()},methods:{ContainerOperator:function(t,e){var n=this;this.listLoading=!0;var r={ContainerId:e.ContainerId,ServerNames:[e.ServerName]};Object(i["a"])(t,r).then((function(t){"100200"===t.code?n.$message({message:n.$t("命令下发成功"),type:"success"}):n.$message({message:t.msg,type:"warning"}),n.listLoading=!1}))},fetchServerNames:function(){var t=this;this.listLoading=!0,Object(o["c"])().then((function(e){t.res.serverNames=e.data,t.listLoading=!1}))},fetchContainerInfoData:function(){var t=this;this.loading=!0,Object(i["c"])().then((function(e){var n=e.data;for(var r in t.res.ContainerNames=[],n)for(var a in n[r].containers){var i=n[r].containers[a].Name;-1===t.res.ContainerNames.indexOf(i)&&t.res.ContainerNames.push(i)}t.loading=!1}))},fetchData:function(){var t=this;this.listLoading=!0,Object(i["d"])(this.listQuery).then((function(e){t.list=e.data,t.list.sort((function(t,e){return t.ServerName.localeCompare(e.ServerName)})),t.groupList=new Set(t.list.map((function(t){return t.ServerName}))),t.groups={};for(var n=0;n<t.list.length;n++){var r=t.list[n],a=t.groups[r.ServerName];a||(a={ServerName:r.ServerName,Members:0,StartIdx:n}),a.Members=a.Members+1,t.groups[r.ServerName]=a}t.listLoading=!1}))},spanMethod:function(t){var e=t.row,n=t.rowIndex,r=t.columnIndex;if(1===r){var a=this.groups[e.ServerName];return n===a.StartIdx?{rowspan:a.Members,colspan:1}:{rowspan:0,colspan:0}}return{rowspan:1,colspan:1}},formatDate:function(t){return Object(s["a"])(t)},PortsToStr:function(t){return Object(l["c"])(t)},PortToStr:function(t){var e=(t.IP?t.IP:"")+(t.PublicPort?":"+t.PublicPort:"");return(e?e+" -> ":"")+t.PrivatePort+"/"+t.Type},FormatName:function(t){return Object(l["a"])(t)},openDetail:function(t){var e=this;this.selectRow=t,this.dialogDetailVisible=!0;var n={ContainerId:t.ContainerId};Object(i["b"])(n).then((function(t){e.selectRow=t.data}))}}},c=u,d=n("2877"),f=Object(d["a"])(c,r,a,!1,null,null,null);e["default"]=f.exports},aa22:function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return i})),n.d(e,"a",(function(){return o}));var r=n("b775");function a(t){return Object(r["a"])({url:"/mgr/server/list",method:"get",params:t})}function i(t){return Object(r["a"])({url:"/mgr/server/names",method:"get",params:t})}function o(t){return Object(r["a"])({url:"/mgr/server/detail",method:"get",params:t})}},b1cf:function(t,e,n){"use strict";n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return a})),n.d(e,"b",(function(){return i})),n.d(e,"d",(function(){return o}));n("b680");function r(t){if(!t||0===t.length)return"";var e="";for(var n in t){var r=t[n],a=(r.IP?r.IP:"")+(r.PublicPort?":"+r.PublicPort:"");e+=", "+(a?a+"->":"")+r.PrivatePort+"/"+r.Type}return e.substr(1,e.length)}function a(t){return t[0].substr(1,t[0].length-1)}function i(t){if(!t||0===t.length)return"";var e=[];for(var n in t){var r=t[n];e.push(r.Source+" : "+r.Destination)}return e}function o(t){if(!t)return"";var e="";return e=t<1024?t+"B":t<1048576?(t/1024).toFixed(3)+"KB":t<1073741824?(t/1048576).toFixed(3)+"MB":(t/1073741824).toFixed(3)+"GB",e}},ed08:function(t,e,n){"use strict";n.d(e,"a",(function(){return r}));n("53ca"),n("ac1f"),n("5319"),n("4d63"),n("25f0"),n("d3b7"),n("4d90"),n("a15b"),n("d81d"),n("b64b"),n("1276"),n("159b");function r(t){if(!t)return"";t=t>0xee09da7916?t:1e3*t;var e=new Date(t),n=e.getFullYear()+"-",r=(e.getMonth()+1<10?"0"+(e.getMonth()+1):e.getMonth()+1)+"-",a=e.getDate()<10?"0"+e.getDate():e.getDate(),i=(e.getHours()<10?"0"+e.getHours():e.getHours())+":",o=(e.getMinutes()<10?"0"+e.getMinutes():e.getMinutes())+":",s=(e.getSeconds()<10?"0"+e.getSeconds():e.getSeconds())+".",l="";return l=e.getMilliseconds()<10?"00"+e.getMilliseconds():e.getMilliseconds()<100?"0"+e.getMilliseconds():e.getMilliseconds(),n+r+a+" "+i+o+s+l}}}]);