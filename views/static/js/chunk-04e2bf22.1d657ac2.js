(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-04e2bf22"],{"20d2":function(e,t,a){"use strict";a("c1cb")},c1cb:function(e,t,a){},e382:function(e,t,a){"use strict";a.r(t);var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"app-container"},[a("div",{staticStyle:{"padding-bottom":"10px"}},[a("el-button",{attrs:{type:"primary"},on:{click:function(t){e.dialogVisible=!0}}},[e._v(e._s(e.$t("添加用户")))])],1),a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.listLoading,expression:"listLoading"}],attrs:{data:e.list,"element-loading-text":"Loading",border:"",fit:"","highlight-current-row":""}},[a("el-table-column",{attrs:{label:"ID",width:"95",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.id)+" ")]}}])}),a("el-table-column",{attrs:{label:"Username",width:"150",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.username)+" ")]}}])}),a("el-table-column",{attrs:{label:"Mobile",width:"150",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.mobile)+" ")]}}])}),a("el-table-column",{attrs:{label:"Email",width:"170",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.email)+" ")]}}])}),a("el-table-column",{attrs:{label:"Nickname",width:"170",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.nickname)+" ")]}}])}),a("el-table-column",{attrs:{label:"Role",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(e.formatRole(t.row.role))+" ")]}}])}),a("el-table-column",{attrs:{label:e.$t("状态"),align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(e.formatStatus(t.row.status))+" ")]}}])}),a("el-table-column",{attrs:{label:"Create",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(e.formatDate(t.row.create_date))+" ")]}}])}),a("el-table-column",{attrs:{label:"Update",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(e.formatDate(t.row.update_date))+" ")]}}])}),a("el-table-column",{attrs:{label:e.$t("操作"),align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-button",{attrs:{size:"small",type:"text"},on:{click:function(a){return e.deleteUser(t.row)}}},[e._v(e._s(e.$t("删除")))]),a("el-button",{attrs:{size:"small",type:"text"},on:{click:function(a){return e.changeStatus(t.row,0===t.row.status?1:0)}}},[e._v(e._s(0===t.row.status?e.$t("禁用"):e.$t("恢复")))]),a("el-button",{attrs:{size:"small",type:"text"},on:{click:function(a){return e.resetPassword(t.row)}}},[e._v(e._s(e.$t("密码重置")))])]}}])})],1),a("el-pagination",{staticStyle:{width:"500px",margin:"0 auto","margin-top":"10px"},attrs:{"hide-on-single-page":!0,"current-page":e.page.currentPage,"page-sizes":[10,30,50,100,200,300,400],"page-size":e.page.pageSize,layout:"prev, pager, next, jumper, sizes, total",total:e.page.total},on:{"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange}}),a("el-dialog",{attrs:{title:e.$t("添加用户"),visible:e.dialogVisible,width:"30%"},on:{"update:visible":function(t){e.dialogVisible=t}}},[a("el-form",{directives:[{name:"loading",rawName:"v-loading",value:e.formLoading,expression:"formLoading"}],attrs:{model:e.formUser,"label-position":"right","label-width":"120px"}},[a("el-form-item",{attrs:{label:e.$t("昵称"),rules:e.rules.nickname}},[a("el-input",{attrs:{placeholder:e.$t("请输入昵称")},model:{value:e.formUser.nickname,callback:function(t){e.$set(e.formUser,"nickname",t)},expression:"formUser.nickname"}})],1),a("el-form-item",{attrs:{label:e.$t("用户名"),rules:e.rules.username}},[a("el-input",{attrs:{placeholder:e.$t("请输入用户名")},model:{value:e.formUser.username,callback:function(t){e.$set(e.formUser,"username",t)},expression:"formUser.username"}})],1),a("el-form-item",{attrs:{label:e.$t("手机"),prop:"mobile",rules:e.rules.mobile}},[a("el-input",{attrs:{placeholder:e.$t("请输入手机号码")},model:{value:e.formUser.mobile,callback:function(t){e.$set(e.formUser,"mobile",t)},expression:"formUser.mobile"}})],1),a("el-form-item",{attrs:{label:e.$t("邮箱"),prop:"email",rules:e.rules.email}},[a("el-input",{attrs:{placeholder:e.$t("请输入邮箱地址")},model:{value:e.formUser.email,callback:function(t){e.$set(e.formUser,"email",t)},expression:"formUser.email"}})],1),a("el-form-item",{attrs:{label:e.$t("密码"),rules:e.rules.password}},[a("el-input",{attrs:{placeholder:e.$t("请输入密码"),"show-password":""},model:{value:e.formUser.password,callback:function(t){e.$set(e.formUser,"password",t)},expression:"formUser.password"}})],1),a("el-form-item",{attrs:{label:e.$t("角色")}},[a("el-select",{attrs:{placeholder:e.$t("请选择")},model:{value:e.formUser.role,callback:function(t){e.$set(e.formUser,"role",t)},expression:"formUser.role"}},e._l(e.roles,(function(e){return a("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1)],1),a("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(t){e.dialogVisible=!1}}},[e._v(e._s(e.$t("取消")))]),a("el-button",{attrs:{type:"primary"},on:{click:e.addUser}},[e._v(e._s(e.$t("确定")))])],1)],1)],1)},l=[],n=a("c24f"),s=a("ed08"),i=a("61f7"),o={components:{},data:function(){return{dialogVisible:!1,list:[],listLoading:!1,roleStr:{ADMIN:this.$t("超级管理员"),AGENT:this.$t("客户端代理")},roles:[{label:this.$t("超级管理员"),value:"ADMIN"},{label:this.$t("客户端代理"),value:"AGENT"}],formLoading:!1,rules:{nickname:[{required:!0,message:this.$t("请输入昵称"),trigger:"blur"},{min:3,max:15,message:this.$t("请输入正确的昵称"),trigger:"blur"}],username:[{required:!0,message:this.$t("请输入用户名"),trigger:"blur"},{min:3,max:20,message:this.$t("请输入正确的用户名"),trigger:"blur"}],checkPhone:[{required:!0,validator:i["b"],trigger:"blur"}],checkId:[{required:!0,validator:i["a"],trigger:"blur"}],mobile:[{required:!0,message:this.$t("请输入手机号码"),trigger:"blur"},{min:6,max:11,message:this.$t("请输入正确的手机号码"),trigger:"blur"},{pattern:/^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$/,message:this.$t("请输入正确的手机号码")}],email:[{required:!0,message:this.$t("请输入邮箱地址"),trigger:"blur"},{type:"email",message:this.$t("请输入正确的邮箱地址"),trigger:["blur","change"]}],password:[{required:!0,message:this.$t("请输入密码"),trigger:"blur"}]},formUser:{nickname:"",username:"",mobile:"",email:"",password:"",role:""},page:{currentPage:1,pageSize:10,total:0}}},created:function(){this.fetchData()},methods:{fetchData:function(){var e=this;this.listLoading=!0,Object(n["g"])(this.page).then((function(t){e.list=t.data.list,e.page=t.data.page,e.listLoading=!1}))},handleSizeChange:function(e){this.page.pageSize=e,this.fetchData()},handleCurrentChange:function(e){this.page.currentPage=e,this.fetchData()},formatDate:function(e){return Object(s["c"])(e)},formatStatus:function(e){return 0===e?this.$t("正常"):this.$t("禁用")},formatRole:function(e){var t={ADMIN:this.$t("超级管理员"),AGENT:this.$t("客户端代理")};return t[e]},changeStatus:function(e,t){var a=this;e.Loading=!0;var r={username:e.username,status:t};Object(n["c"])(r).then((function(e){e.Loading=!1,a.fetchData()}))},resetPassword:function(e){var t=this;e.Loading=!0;var a={username:e.username};Object(n["j"])(a).then((function(e){e.Loading=!1,t.$alert(t.$t("密码仅显示一次，请备份：")+e.data,t.$t("密码重置成功"),{confirmButtonText:t.$t("确定")}),t.fetchData()}))},deleteUser:function(e){var t=this;this.$confirm(this.$t("此操作将永久删除该用户, 是否继续?"),this.$t("提示"),{confirmButtonText:this.$t("确定"),cancelButtonText:this.$t("取消"),type:"warning"}).then((function(){e.Loading=!0;var a={id:e.id};Object(n["d"])(a).then((function(e){e.Loading=!1,t.fetchData()}))}))},addUser:function(){var e=this;this.formLoading=!0,Object(n["a"])(this.formUser).then((function(t){e.formLoading=!1,e.dialogVisible=!1,e.fetchData()})).catch((function(t){console.log(t),e.formLoading=!1}))}}},u=o,c=(a("20d2"),a("2877")),m=Object(c["a"])(u,r,l,!1,null,"1cc717a3",null);t["default"]=m.exports}}]);