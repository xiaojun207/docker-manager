import{_ as k,Y as y,H as C,r,x as E,o as a,c as f,y as M,z as s,w as l,b as o,e as c,t as _,A as p,a as $}from"./index-18b9e6d0.js";import{g as x,u as D}from"./config-99ec0a22.js";const L={data(){return{listLoading:!1,list:[],ElIconEdit:y,ElIconRefresh:C}},created(){this.fetchData()},methods:{fetchData(){this.listLoading=!0,x().then(t=>{this.list=t.data.map(i=>(i.edit=!1,i.originalValue=i.Value,i.originalMemo=i.Memo,i)),this.listLoading=!1})},cancelEdit(t){t.Value=t.originalValue,t.Memo=t.originalMemo,t.edit=!1},confirmEdit(t){t.edit=!1,t.originalValue=t.Value,t.originalMemo=t.Memo,this.listLoading=!0,D(t).then(i=>{this.$message(this.$t("发布成功")),this.listLoading=!1,this.fetchData()})}}},z={class:"app-container"},v=$("div",{style:{"margin-bottom":"15px"}},null,-1),I={key:1},N={key:1};function B(t,i,U,R,u,h){const d=r("el-table-column"),m=r("el-button"),g=r("el-input"),V=r("el-table"),b=E("loading");return a(),f("div",z,[v,M((a(),s(V,{data:u.list,"element-loading-text":"Loading",border:"",fit:"","highlight-current-row":""},{default:l(()=>[o(d,{align:"center",label:"ID",width:"95"},{default:l(e=>[c(_(e.$index+1),1)]),_:1}),o(d,{label:t.$t("字段名称"),width:"270"},{default:l(e=>[o(m,{type:"primary",link:""},{default:l(()=>[c(_(e.row.Name),1)]),_:2},1024)]),_:1},8,["label"]),o(d,{"min-width":"270px",label:t.$t("值")},{default:l(({row:e})=>[e.edit?(a(),s(g,{key:0,modelValue:e.Value,"onUpdate:modelValue":n=>e.Value=n,class:"edit-input",size:"small"},null,8,["modelValue","onUpdate:modelValue"])):(a(),f("span",I,_(e.Value),1))]),_:1},8,["label"]),o(d,{label:t.$t("备注"),width:"270"},{default:l(({row:e})=>[e.edit?(a(),s(g,{key:0,modelValue:e.Memo,"onUpdate:modelValue":n=>e.Memo=n,class:"edit-input",size:"small"},null,8,["modelValue","onUpdate:modelValue"])):(a(),f("span",N,_(e.Memo),1))]),_:1},8,["label"]),o(d,{align:"center",label:t.$t("操作"),width:"270"},{default:l(({row:e})=>[e.edit?p("",!0):(a(),s(m,{key:0,type:"primary",size:"small",icon:u.ElIconEdit,onClick:n=>e.edit=!e.edit},{default:l(()=>[c("Edit")]),_:2},1032,["icon","onClick"])),e.edit?(a(),s(m,{key:1,type:"success",size:"small",onClick:n=>h.confirmEdit(e)},{default:l(()=>[c("OK")]),_:2},1032,["onClick"])):p("",!0),e.edit?(a(),s(m,{key:2,type:"warning",size:"small",icon:u.ElIconRefresh,onClick:n=>h.cancelEdit(e)},{default:l(()=>[c("Cancel")]),_:2},1032,["icon","onClick"])):p("",!0)]),_:1},8,["label"])]),_:1},8,["data"])),[[b,u.listLoading]])])}const K=k(L,[["render",B]]);export{K as default};
