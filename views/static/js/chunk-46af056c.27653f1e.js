(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-46af056c"],{"4dd0":function(e,t,r){"use strict";r.d(t,"d",(function(){return a})),r.d(t,"b",(function(){return i})),r.d(t,"c",(function(){return s})),r.d(t,"f",(function(){return o})),r.d(t,"a",(function(){return l}));var n=r("b775");function a(e){return Object(n["a"])({url:"/mgr/container/list",method:"get",params:e})}function i(e){return Object(n["a"])({url:"/mgr/container/detail",method:"get",params:e})}function s(e){return Object(n["a"])({url:"/mgr/containerInfos",method:"get",params:e})}function o(e){return Object(n["a"])({url:"/mgr/publish",method:"post",data:e})}function l(e,t){return Object(n["a"])({url:"/mgr/container/"+e,method:"post",data:t})}},b1cf:function(e,t,r){"use strict";r.d(t,"c",(function(){return n})),r.d(t,"a",(function(){return a})),r.d(t,"b",(function(){return i})),r.d(t,"d",(function(){return s}));r("b680");function n(e){if(!e||0===e.length)return"";var t="";for(var r in e){var n=e[r],a=(n.IP?n.IP:"")+(n.PublicPort?":"+n.PublicPort:"");t+=", "+(a?a+"->":"")+n.PrivatePort+"/"+n.Type}return t.substr(1,t.length)}function a(e){return e[0].substr(1,e[0].length-1)}function i(e){if(!e||0===e.length)return"";var t=[];for(var r in e){var n=e[r];t.push(n.Source+" : "+n.Destination)}return t}function s(e){if(!e)return"";var t="";return t=e<1024?e+"B":e<1048576?(e/1024).toFixed(3)+"KB":e<1073741824?(e/1048576).toFixed(3)+"MB":(e/1073741824).toFixed(3)+"GB",t}},e036:function(e,t,r){"use strict";r.r(t);var n=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"app-container"},[r("div",{staticClass:"filter-container"},[r("el-select",{staticClass:"filter-item",staticStyle:{width:"300px","margin-right":"10px"},attrs:{multiple:"",filterable:"",placeholder:e.$t("服务器"),clearable:"","collapse-tags":""},model:{value:e.listQuery.serverNames,callback:function(t){e.$set(e.listQuery,"serverNames",t)},expression:"listQuery.serverNames"}},e._l(e.res.serverNames,(function(e){return r("el-option",{key:e,attrs:{label:e,value:e}})})),1),r("el-button",{staticClass:"filter-item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:e.fetchData}},[e._v(" Search ")])],1),r("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.listLoading,expression:"listLoading"}],ref:"tableList",staticClass:"filter-tree",staticStyle:{"margin-top":"20px"},attrs:{data:e.list.filter((function(t){return!e.search||t.RepoTags.toLowerCase().includes(e.search.toLowerCase())})),"span-method":e.spanMethod,"element-loading-text":"Loading",stripe:"",border:"",fit:"","highlight-current-row":""}},[r("el-table-column",{attrs:{align:"center",label:"ID",width:"65"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.$index+1)+" ")]}}])}),r("el-table-column",{attrs:{label:"ServerName",width:"170"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.ServerName)+" ")]}}])}),r("el-table-column",{attrs:{label:"ImageId",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("el-tooltip",{staticClass:"item",attrs:{effect:"dark",content:t.row.ImageId,placement:"top-start"}},[r("el-button",{attrs:{type:"text"},on:{click:function(r){return e.openDetail(t.row)}}},[e._v(e._s(e.formatImageId(t.row.ImageId)))])],1)]}}])}),r("el-table-column",{attrs:{label:"Tags",align:"center"},scopedSlots:e._u([{key:"header",fn:function(t){return[e._v(" Tags "),r("el-input",{staticStyle:{width:"140px"},attrs:{size:"mini",placeholder:e.$t("输入关键字过滤")},model:{value:e.search,callback:function(t){e.search=t},expression:"search"}})]}},{key:"default",fn:function(t){return e._l(e.formatTags(t.row.RepoTags),(function(t){return r("div",{key:t},[e._v(e._s(t))])}))}}])}),r("el-table-column",{attrs:{label:"Size",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(e.formatSize(t.row.Size))+" ")]}}])}),r("el-table-column",{attrs:{label:"RepoDigests",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.RepoDigests)+" ")]}}])}),r("el-table-column",{attrs:{align:"center",prop:"created_at",label:"Created",width:"200"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("i",{staticClass:"el-icon-time"}),r("span",[e._v(e._s(e.formatDate(t.row.Created)))])]}}])}),r("el-table-column",{attrs:{"class-name":"status-col",label:e.$t("操作"),width:"170",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("el-button",{attrs:{loading:e.listLoading,type:"text"},on:{click:function(r){return e.ImageOperator("remove",t.row)}}},[e._v(e._s(e.$t("删除")))])]}}])})],1),r("el-dialog",{directives:[{name:"loading",rawName:"v-loading",value:e.detailLoading,expression:"detailLoading"}],attrs:{visible:e.dialogDetailVisible,title:e.$t("详情")},on:{"update:visible":function(t){e.dialogDetailVisible=t}}},[r("pre",[e._v(e._s(JSON.stringify(e.selectRow,null,2))+"\n      ")])])],1)},a=[],i=(r("d3b7"),r("6062"),r("3ca3"),r("ddb0"),r("d81d"),r("ac1f"),r("1276"),r("4dd0")),s=r("b775");function o(e){return Object(s["a"])({url:"/mgr/image/list",method:"get",params:e})}function l(e,t){return Object(s["a"])({url:"/mgr/image/"+e,method:"post",data:t})}function u(e){return console.log("getImage.params:",e),Object(s["a"])({url:"/mgr/image/detail",method:"get",params:e})}var c=r("ed08"),d=r("b1cf"),f={filters:{statusFilter:function(e){var t={published:"success",draft:"gray",deleted:"danger"};return t[e]}},data:function(){return{list:[],groupList:[],groups:{},listLoading:!1,detailLoading:!1,dialogDetailVisible:!1,selectRow:{},res:{serverNames:[]},listQuery:{serverNames:[]},filterText:"",search:""}},computed:{},created:function(){this.fetchData(),this.fetchServerNames()},methods:{ImageOperator:function(e,t){var r=this;this.listLoading=!0;var n={ImageId:t.ImageId,serverName:t.ServerName};l(e,n).then((function(e){"100200"===e.code&&(r.$message({message:r.$t("命令下发成功"),type:"success"}),r.fetchData()),r.listLoading=!1}))},fetchServerNames:function(){var e=this;Object(i["getServerNames"])().then((function(t){e.res.serverNames=t.data}))},fetchData:function(){var e=this;this.listLoading=!0,o(this.listQuery).then((function(t){e.list=t.data,e.list.sort((function(e,t){return e.ServerName.localeCompare(t.ServerName)})),e.groupList=new Set(e.list.map((function(e){return e.ServerName}))),e.groups={};for(var r=0;r<e.list.length;r++){var n=e.list[r],a=e.groups[n.ServerName];a||(a={ServerName:n.ServerName,Members:0,StartIdx:r}),a.Members=a.Members+1,e.groups[n.ServerName]=a}e.listLoading=!1}))},spanMethod:function(e){var t=e.row,r=e.rowIndex,n=e.columnIndex;if(1===n){var a=this.groups[t.ServerName];return r===a.StartIdx?{rowspan:a.Members,colspan:1}:{rowspan:0,colspan:0}}return{rowspan:1,colspan:1}},formatDate:function(e){return Object(c["a"])(e)},FormatName:function(e){return Object(d["a"])(e)},formatSize:function(e){return Object(d["d"])(e)},formatTags:function(e){return e.split(",")},formatImageId:function(e){return e.substr(7,12)},openDetail:function(e){var t=this;this.dialogDetailVisible=!0,this.detailLoading=!0,this.selectRow=e;var r={ImageId:e.ImageId};u(r).then((function(e){t.detailLoading=!1,t.selectRow=e.data,console.log(e.data)}))}}},g=f,m=r("2877"),p=Object(m["a"])(g,n,a,!1,null,null,null);t["default"]=p.exports},ed08:function(e,t,r){"use strict";r.d(t,"a",(function(){return n}));r("53ca"),r("ac1f"),r("5319"),r("4d63"),r("25f0"),r("d3b7"),r("4d90"),r("a15b"),r("d81d"),r("b64b"),r("1276"),r("159b");function n(e){if(!e)return"";e=e>0xee09da7916?e:1e3*e;var t=new Date(e),r=t.getFullYear()+"-",n=(t.getMonth()+1<10?"0"+(t.getMonth()+1):t.getMonth()+1)+"-",a=t.getDate()<10?"0"+t.getDate():t.getDate(),i=(t.getHours()<10?"0"+t.getHours():t.getHours())+":",s=(t.getMinutes()<10?"0"+t.getMinutes():t.getMinutes())+":",o=(t.getSeconds()<10?"0"+t.getSeconds():t.getSeconds())+".",l="";return l=t.getMilliseconds()<10?"00"+t.getMilliseconds():t.getMilliseconds()<100?"0"+t.getMilliseconds():t.getMilliseconds(),r+n+a+" "+i+s+o+l}}}]);