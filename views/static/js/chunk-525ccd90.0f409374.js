(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-525ccd90"],{9410:function(t,e,n){},d5db:function(t,e,n){"use strict";n("9410")},eb43:function(t,e,n){"use strict";n.r(e);var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"app-container"},[n("div",{staticStyle:{"margin-bottom":"15px"}}),n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.listLoading,expression:"listLoading"}],attrs:{data:t.list,"element-loading-text":"Loading",border:"",fit:"","highlight-current-row":""}},[n("el-table-column",{attrs:{align:"center",label:"ID",width:"95"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.$index+1)+" ")]}}])}),n("el-table-column",{attrs:{label:t.$t("字段名称"),width:"270"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-button",{attrs:{type:"text"}},[t._v(t._s(e.row.Name))])]}}])}),n("el-table-column",{attrs:{label:t.$t("值"),align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.Value)+" ")]}}])}),n("el-table-column",{attrs:{label:t.$t("备注"),width:"270"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.Memo)+" ")]}}])})],1)],1)},i=[],o=n("b775");function l(){return Object(o["a"])({url:"/mgr/config",method:"get"})}function s(t){return Object(o["a"])({url:"/mgr/config/update",method:"post",data:t})}var r={data:function(){return{list:[],form:{TaskFrequency:600}}},created:function(){this.fetchData()},methods:{onSubmit:function(){var t=this;this.listLoading=!0,console.log("this.form:",this.form),s(this.form).then((function(e){t.$message(t.$t("发布成功")),t.listLoading=!1}))},fetchData:function(){var t=this;this.listLoading=!0,l().then((function(e){t.list=e.data,t.listLoading=!1}))}}},u=r,c=(n("d5db"),n("2877")),d=Object(c["a"])(u,a,i,!1,null,"b32823bc",null);e["default"]=d.exports}}]);