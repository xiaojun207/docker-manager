import{a as z,u as U,T as G,P as j,S as I,D as O,b as J,d as M}from"./table.00d3b251.js";import{d as L,L as P,k as y,_ as N,r as u,o as b,v,w as n,b as o,c as T,x as A,F as S,e as $,t as B,j as E,A as K,a as _,B as H,C as Q,E as R}from"./index.3f48ed1b.js";function W(e){const a=e,r=e/1024,p=e/(1024*1024),s=e/(1024*1024*1024);return s>.1?s.toFixed(4)+"GB":p>.1?p.toFixed(4)+"MB":r>.1?r.toFixed(4)+"KB":a.toFixed(4)+"b"}const k=[{value:1,label:"\u8FD0\u52A8"},{value:2,label:"\u5065\u8EAB"},{value:3,label:"\u8DD1\u9177"},{value:4,label:"\u8857\u821E"}],q=[{value:1,label:"\u4ECA\u5929"},{value:2,label:"\u660E\u5929"},{value:3,label:"\u540E\u5929"}],X=L({components:{Layer:P},props:{layer:{type:Object,default:()=>({show:!1,title:"",showButton:!0})}},setup(e,a){const r=y(null),p=y(null);let s=y({name:"",choose:"",radio:"",number:""});const h={name:[{required:!0,message:"\u8BF7\u8F93\u5165\u59D3\u540D",trigger:"blur"}],number:[{required:!0,message:"\u8BF7\u8F93\u5165\u6570\u5B57",trigger:"blur"}],choose:[{required:!0,message:"\u8BF7\u9009\u62E9",trigger:"blur"}],radio:[{required:!0,message:"\u8BF7\u9009\u62E9",trigger:"blur"}]};c();function c(){e.layer.row&&(s.value=JSON.parse(JSON.stringify(e.layer.row)))}return{form:s,rules:h,layerDom:p,ruleForm:r,selectData:k,radioData:q}},methods:{submit(){this.ruleForm&&this.ruleForm.validate(e=>{if(e){let a=this.form;this.layer.row?this.updateForm(a):this.addForm(a)}else return!1})},addForm(e){z(e).then(a=>{this.$message({type:"success",message:"\u65B0\u589E\u6210\u529F"}),this.$emit("getTableData",!0),this.layerDom&&this.layerDom.close()})},updateForm(e){U(e).then(a=>{this.$message({type:"success",message:"\u7F16\u8F91\u6210\u529F"}),this.$emit("getTableData",!1),this.layerDom&&this.layerDom.close()})}}});function Y(e,a,r,p,s,h){const c=u("el-input"),i=u("el-form-item"),F=u("el-option"),f=u("el-select"),D=u("el-radio"),m=u("el-radio-group"),g=u("el-form"),t=u("Layer",!0);return b(),v(t,{layer:e.layer,onConfirm:e.submit,ref:"layerDom"},{default:n(()=>[o(g,{model:e.form,rules:e.rules,ref:"ruleForm","label-width":"120px",style:{"margin-right":"30px"}},{default:n(()=>[o(i,{label:"\u540D\u79F0\uFF1A",prop:"name"},{default:n(()=>[o(c,{modelValue:e.form.name,"onUpdate:modelValue":a[0]||(a[0]=l=>e.form.name=l),placeholder:"\u8BF7\u8F93\u5165\u540D\u79F0"},null,8,["modelValue"])]),_:1}),o(i,{label:"\u6570\u5B57\uFF1A",prop:"number"},{default:n(()=>[o(c,{modelValue:e.form.number,"onUpdate:modelValue":a[1]||(a[1]=l=>e.form.number=l),oninput:"value=value.replace(/[^\\d]/g,'')",placeholder:"\u53EA\u80FD\u8F93\u5165\u6B63\u6574\u6570"},null,8,["modelValue"])]),_:1}),o(i,{label:"\u9009\u62E9\u5668\uFF1A",prop:"select"},{default:n(()=>[o(f,{modelValue:e.form.choose,"onUpdate:modelValue":a[2]||(a[2]=l=>e.form.choose=l),placeholder:"\u8BF7\u9009\u62E9",clearable:""},{default:n(()=>[(b(!0),T(S,null,A(e.selectData,l=>(b(),v(F,{key:l.value,label:l.label,value:l.value},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1}),o(i,{label:"\u5355\u9009\u6846\uFF1A",prop:"radio"},{default:n(()=>[o(m,{modelValue:e.form.radio,"onUpdate:modelValue":a[3]||(a[3]=l=>e.form.radio=l)},{default:n(()=>[(b(!0),T(S,null,A(e.radioData,l=>(b(),v(D,{key:l.value,label:l.value},{default:n(()=>[$(B(l.label),1)]),_:2},1032,["label"]))),128))]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["model","rules"])]),_:1},8,["layer","onConfirm"])}var Z=N(X,[["render",Y]]);const x=L({name:"crudTable",components:{Table:G,Layer:Z},setup(){const e=E({name:""}),a=E({show:!1,title:"\u65B0\u589E",showButton:!0}),r=E({index:1,size:20,total:0}),p=y(!0),s=y([]),h=y([]),c=m=>{h.value=m},i=m=>{p.value=!0,m&&(r.index=1);let g={page_num:r.index,page_size:r.size,...e};J(g).then(t=>{let l=t.data.list;Array.isArray(l)&&l.forEach(d=>{d.update_time=new Date(d.update_time),d.size=W(d.size);const C=k.find(w=>w.value===d.choose);C?d.chooseName=C.label:d.chooseName=d.choose;const V=q.find(w=>w.value===d.radio);V?d.radioName=V.label:d.radio}),s.value=l,r.total=Number(t.data.page.total)}).catch(t=>{s.value=[],r.index=1,r.total=0}).finally(()=>{p.value=!1})},F=m=>{let g={names:m.map(t=>t.name).join(",")};M(g).then(t=>{R({type:"success",message:"\u5220\u9664\u6210\u529F"}),i(s.value.length===1)})},f=()=>{a.title="\u65B0\u589E\u6570\u636E",a.show=!0,delete a.row},D=m=>{a.title="\u7F16\u8F91\u6570\u636E",a.row=m,a.show=!0};return i(!0),{Plus:j,Search:I,Delete:O,query:e,tableData:s,chooseData:h,loading:p,page:r,layer:a,handleSelectionChange:c,handleAdd:f,handleEdit:D,handleDel:F,getTableData:i}}}),ee={class:"layout-container"},ae={class:"layout-container-form flex space-between"},le={class:"layout-container-form-handle"},oe={class:"layout-container-form-search"},te={class:"layout-container-table"};function ne(e,a,r,p,s,h){const c=u("el-button"),i=u("el-popconfirm"),F=u("el-input"),f=u("el-table-column"),D=u("Table"),m=u("Layer"),g=K("loading");return b(),T("div",ee,[_("div",ae,[_("div",le,[o(i,{title:e.$t("message.common.delTip"),onConfirm:a[0]||(a[0]=t=>e.handleDel(e.chooseData))},{reference:n(()=>[o(c,{type:"danger",icon:e.Delete,disabled:e.chooseData.length===0},{default:n(()=>[$(B(e.$t("message.common.delBat")),1)]),_:1},8,["icon","disabled"])]),_:1},8,["title"])]),_("div",oe,[o(F,{modelValue:e.query.name,"onUpdate:modelValue":a[1]||(a[1]=t=>e.query.name=t),placeholder:e.$t("message.common.searchTip"),onChange:a[2]||(a[2]=t=>e.getTableData(!0))},null,8,["modelValue","placeholder"]),o(c,{type:"primary",icon:e.Search,class:"search-btn",onClick:a[3]||(a[3]=t=>e.getTableData(!0))},{default:n(()=>[$(B(e.$t("message.common.search")),1)]),_:1},8,["icon"])])]),_("div",te,[H((b(),v(D,{ref:"table",page:e.page,"onUpdate:page":a[4]||(a[4]=t=>e.page=t),showIndex:!0,showSelection:!0,data:e.tableData,onGetTableData:e.getTableData,onSelectionChange:e.handleSelectionChange},{default:n(()=>[o(f,{prop:"name",label:"\u540D\u79F0",align:"center"}),o(f,{prop:"size",label:"\u5927\u5C0F",align:"center"}),o(f,{prop:"num",label:"\u6570\u91CF",align:"center"}),o(f,{prop:"update_time",label:"\u66F4\u65B0\u65F6\u95F4",align:"center"}),o(f,{label:e.$t("message.common.handle"),align:"center",fixed:"right",width:"200"},{default:n(t=>[o(i,{title:e.$t("message.common.delTip"),onConfirm:l=>e.handleDel([t.row])},{reference:n(()=>[o(c,{type:"danger"},{default:n(()=>[$(B(e.$t("message.common.del")),1)]),_:1})]),_:2},1032,["title","onConfirm"])]),_:1},8,["label"])]),_:1},8,["page","data","onGetTableData","onSelectionChange"])),[[g,e.loading]]),e.layer.show?(b(),v(m,{key:0,layer:e.layer,onGetTableData:e.getTableData},null,8,["layer","onGetTableData"])):Q("",!0)])])}var se=N(x,[["render",ne]]);export{se as default};