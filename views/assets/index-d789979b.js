import{q as v,_ as k,R,L as M,r as g,x as T,o as u,c as h,a as _,b as s,w as r,J as I,K as w,z as S,e as m,y as N,t as c}from"./index-94dbf9cf.js";import{g as Q}from"./server-13ad57c4.js";import{a as P,f as B}from"./docker-6b00f3d6.js";function F(e){return v({url:"/mgr/image/list",method:"get",params:e})}function O(e,a){return v({url:"/mgr/image/"+e,method:"post",data:a})}function E(e){return console.log("getImage.params:",e),v({url:"/mgr/image/detail",method:"get",params:e})}const U={components:{},data(){return{list:[],page:{currentPage:1,pageSize:10,total:0},groupList:[],groups:{},listLoading:!1,detailLoading:!1,dialogDetailVisible:!1,selectRow:{},res:{serverNames:[]},listQuery:{serverNames:[]},filterText:"",search:"",ElIconSearch:R}},computed:{},created(){this.fetchData(),this.fetchServerNames()},methods:{statusFilter_filter(e){return{published:"success",draft:"gray",deleted:"danger"}[e]},ImageOperator(e,a){this.listLoading=!0;const i={ImageId:a.ImageId,serverName:a.ServerName};O(e,i).then(o=>{o.code==="100200"&&(this.$message({message:this.$t("命令下发成功"),type:"success"}),this.fetchData()),this.listLoading=!1})},fetchServerNames(){Q().then(e=>{this.res.serverNames=e.data})},handleSizeChange(e){console.log(`每页 ${e} 条`),this.page.pageSize=e,this.fetchData()},handleCurrentChange(e){console.log(`当前页: ${e}`),this.page.currentPage=e,this.fetchData()},fetchData(){this.listLoading=!0,this.listQuery.currentPage=this.page.currentPage,this.listQuery.pageSize=this.page.pageSize,F(this.listQuery).then(e=>{this.list=e.data.list,this.page=e.data.page,this.list||(this.list=[]),this.list.sort(function(a,i){return a.ServerName.localeCompare(i.ServerName)}),this.groupList=new Set(this.list.map(a=>a.ServerName)),this.groups={};for(let a=0;a<this.list.length;a++){const i=this.list[a];let o=this.groups[i.ServerName];o||(o={ServerName:i.ServerName,Members:0,StartIdx:a}),o.Members=o.Members+1,this.groups[i.ServerName]=o}this.listLoading=!1})},spanMethod(e){const{row:a,rowIndex:i,columnIndex:o}=e;if(o===1){const l=this.groups[a.ServerName];return i===l.StartIdx?{rowspan:l.Members,colspan:1}:{rowspan:0,colspan:0}}else return{rowspan:1,colspan:1}},formatDate(e){return M(e)},FormatName(e){return P(e)},formatSize(e){return B(e)},formatTags(e){return e.split(",")},formatImageId(e){return e.substr(7,12)},openDetail(e){this.dialogDetailVisible=!0,this.detailLoading=!0,this.selectRow=e;const a={ImageId:e.ImageId};E(a).then(i=>{this.detailLoading=!1,this.selectRow=i.data,console.log(i.data)})}}},J={class:"app-container"},j={class:"filter-container"},q={style:{"text-indent":"-2em"}};function K(e,a,i,o,l,n){const C=g("el-option"),y=g("el-select"),f=g("el-button"),d=g("el-table-column"),D=g("el-input"),z=g("el-tooltip"),L=g("el-table"),x=g("el-pagination"),V=g("el-dialog"),b=T("loading");return u(),h("div",J,[_("div",j,[s(y,{modelValue:l.listQuery.serverNames,"onUpdate:modelValue":a[0]||(a[0]=t=>l.listQuery.serverNames=t),multiple:"",filterable:"",placeholder:e.$t("主机"),clearable:"","collapse-tags":"",class:"filter-item",style:{width:"300px","margin-right":"10px"}},{default:r(()=>[(u(!0),h(I,null,w(l.res.serverNames,t=>(u(),S(C,{key:t,label:t,value:t},null,8,["label","value"]))),128))]),_:1},8,["modelValue","placeholder"]),s(f,{class:"filter-item",type:"primary",icon:l.ElIconSearch,onClick:n.fetchData},{default:r(()=>[m(" Search ")]),_:1},8,["icon","onClick"])]),N((u(),S(L,{ref:"tableList",data:l.list.filter(t=>!l.search||t.RepoTags.toLowerCase().includes(l.search.toLowerCase())),"span-method":n.spanMethod,"element-loading-text":"Loading",stripe:"",border:"",fit:"","highlight-current-row":"",class:"filter-tree",style:{"margin-top":"20px"}},{default:r(()=>[s(d,{align:"center",label:"ID",width:"65"},{default:r(t=>[m(c(t.$index+1),1)]),_:1}),s(d,{label:"ServerName",width:"170"},{default:r(t=>[m(c(t.row.ServerName),1)]),_:1}),s(d,{label:"RepoDigests",align:"center"},{default:r(t=>[m(c(t.row.RepoDigests.split("@")[0]),1)]),_:1}),s(d,{label:"Tags",align:"center"},{header:r(()=>[m(" Tags "),s(D,{modelValue:l.search,"onUpdate:modelValue":a[1]||(a[1]=t=>l.search=t),size:"small",placeholder:e.$t("输入关键字过滤"),style:{width:"140px"}},null,8,["modelValue","placeholder"])]),default:r(t=>[(u(!0),h(I,null,w(n.formatTags(t.row.RepoTags),p=>(u(),h("div",{key:p},c(p),1))),128))]),_:1}),s(d,{label:"ImageId",align:"center"},{default:r(t=>[s(z,{class:"item",effect:"dark",content:t.row.ImageId,placement:"top-start"},{default:r(()=>[s(f,{type:"primary",link:"",onClick:p=>n.openDetail(t.row)},{default:r(()=>[m(c(n.formatImageId(t.row.ImageId)),1)]),_:2},1032,["onClick"])]),_:2},1032,["content"])]),_:1}),s(d,{label:"Size",align:"center"},{default:r(t=>[m(c(n.formatSize(t.row.Size)),1)]),_:1}),s(d,{align:"center",prop:"created_at",label:"Created",width:"200"},{default:r(t=>[_("span",null,c(n.formatDate(t.row.Created)),1)]),_:1}),s(d,{"class-name":"status-col",label:e.$t("操作"),width:"170",align:"center"},{default:r(t=>[s(f,{loading:l.listLoading,type:"danger",link:"",onClick:p=>n.ImageOperator("remove",t.row)},{default:r(()=>[m(c(e.$t("删除")),1)]),_:2},1032,["loading","onClick"])]),_:1},8,["label"])]),_:1},8,["data","span-method"])),[[b,l.listLoading]]),s(x,{"hide-on-single-page":!0,style:{width:"500px",margin:"10px auto 0"},"current-page":l.page.currentPage,"page-sizes":[10,30,50,100,200,300,400],"page-size":l.page.pageSize,total:l.page.total,layout:"prev, pager, next, jumper, sizes, total",onSizeChange:n.handleSizeChange,onCurrentChange:n.handleCurrentChange},null,8,["current-page","page-size","total","onSizeChange","onCurrentChange"]),N((u(),S(V,{modelValue:l.dialogDetailVisible,"onUpdate:modelValue":a[2]||(a[2]=t=>l.dialogDetailVisible=t),title:e.$t("详情")},{default:r(()=>[_("pre",q,"    "+c(JSON.stringify(l.selectRow,null,2))+`
          `,1)]),_:1},8,["modelValue","title"])),[[b,l.detailLoading]])])}const W=k(U,[["render",K]]);export{W as default};