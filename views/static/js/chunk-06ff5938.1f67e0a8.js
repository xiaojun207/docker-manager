(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-06ff5938"],{da71:function(t,e,n){"use strict";n.d(e,"c",(function(){return a})),n.d(e,"e",(function(){return l})),n.d(e,"d",(function(){return o})),n.d(e,"b",(function(){return r})),n.d(e,"a",(function(){return c}));var i=n("b775");function a(){return Object(i["a"])({url:"/mgr/config",method:"get"})}function l(t){return Object(i["a"])({url:"/mgr/config/update",method:"post",data:t})}function o(){return Object(i["a"])({url:"/mgr/config/whiteList",method:"get"})}function r(t){return Object(i["a"])({url:"/mgr/config/deleteWhiteIp",method:"post",data:t})}function c(t){return Object(i["a"])({url:"/mgr/config/addWhiteIp",method:"post",data:t})}},eb43:function(t,e,n){"use strict";n.r(e);var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"app-container"},[n("div",{staticStyle:{"margin-bottom":"15px"}}),n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.listLoading,expression:"listLoading"}],attrs:{data:t.list,"element-loading-text":"Loading",border:"",fit:"","highlight-current-row":""}},[n("el-table-column",{attrs:{align:"center",label:"ID",width:"95"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.$index+1)+" ")]}}])}),n("el-table-column",{attrs:{label:t.$t("字段名称"),width:"270"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-button",{attrs:{type:"text"}},[t._v(t._s(e.row.Name))])]}}])}),n("el-table-column",{attrs:{"min-width":"270px",label:t.$t("值")},scopedSlots:t._u([{key:"default",fn:function(e){var i=e.row;return[i.edit?[n("el-input",{staticClass:"edit-input",attrs:{size:"small"},model:{value:i.Value,callback:function(e){t.$set(i,"Value",e)},expression:"row.Value"}})]:n("span",[t._v(t._s(i.Value))])]}}])}),n("el-table-column",{attrs:{label:t.$t("备注"),width:"270"},scopedSlots:t._u([{key:"default",fn:function(e){var i=e.row;return[i.edit?[n("el-input",{staticClass:"edit-input",attrs:{size:"small"},model:{value:i.Memo,callback:function(e){t.$set(i,"Memo",e)},expression:"row.Memo"}})]:n("span",[t._v(t._s(i.Memo))])]}}])}),n("el-table-column",{attrs:{align:"center",label:t.$t("操作"),width:"270"},scopedSlots:t._u([{key:"default",fn:function(e){var i=e.row;return[i.edit?t._e():n("el-button",{attrs:{type:"primary",size:"small",icon:"el-icon-edit"},on:{click:function(t){i.edit=!i.edit}}},[t._v("Edit")]),i.edit?n("el-button",{attrs:{type:"success",size:"small",icon:"el-icon-circle-check-outline"},on:{click:function(e){return t.confirmEdit(i)}}},[t._v("OK")]):t._e(),i.edit?n("el-button",{attrs:{type:"warning",size:"small",icon:"el-icon-refresh"},on:{click:function(e){return t.cancelEdit(i)}}},[t._v("Cancel")]):t._e()]}}])})],1)],1)},a=[],l=(n("d81d"),n("da71")),o={data:function(){return{listLoading:!1,list:[]}},created:function(){this.fetchData()},methods:{fetchData:function(){var t=this;this.listLoading=!0,Object(l["c"])().then((function(e){t.list=e.data.map((function(e){return t.$set(e,"edit",!1),e.originalValue=e.Value,e.originalMemo=e.Memo,e})),t.listLoading=!1}))},cancelEdit:function(t){t.Value=t.originalValue,t.Memo=t.originalMemo,t.edit=!1},confirmEdit:function(t){var e=this;t.edit=!1,t.originalValue=t.Value,t.originalMemo=t.Memo,this.listLoading=!0,Object(l["e"])(t).then((function(t){e.$message(e.$t("发布成功")),e.listLoading=!1,e.fetchData()}))}}},r=o,c=n("2877"),u=Object(c["a"])(r,i,a,!1,null,null,null);e["default"]=u.exports}}]);