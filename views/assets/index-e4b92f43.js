import{_ as D,R as P,L as I,r as c,x as L,o as g,c as f,a as C,b as r,w as n,J as _,K as b,z as v,e as m,y as z,t as o,Q as M}from"./index-18b9e6d0.js";import{C as Q,g as O,b as T,c as R}from"./container-3bd7e1dc.js";import{g as U}from"./server-c2b48054.js";import{P as B,a as E}from"./docker-6b00f3d6.js";const F={components:{},data(){return{list:[],page:{currentPage:1,pageSize:10,total:0},groupList:[],groups:{},listLoading:!1,dialogDetailVisible:!1,selectRow:{},filterSearch:{Name:""},res:{serverNames:[],ContainerNames:[],containerInfos:[]},listQuery:{ServerNames:[],ContainerNames:[],state:void 0},ElIconSearch:P}},computed:{},created(){this.fetchData(),this.fetchContainerInfoData(),this.fetchServerNames()},methods:{statusFilter_filter(e){return{published:"success",draft:"gray",deleted:"danger"}[e]},ContainerOperator(e,a){this.listLoading=!0;const i={ContainerId:a.ContainerId,ServerNames:[a.ServerName]};Q(e,i).then(s=>{s.code==="100200"?this.$message({message:this.$t("命令下发成功"),type:"success"}):this.$message({message:s.msg,type:"warning"}),this.listLoading=!1})},fetchServerNames(){this.listLoading=!0,U().then(e=>{this.res.serverNames=e.data,this.listLoading=!1})},fetchContainerInfoData(){this.loading=!0,O().then(e=>{const a=e.data;this.res.ContainerNames=[];for(const i in a)for(const s in a[i].containers){const l=a[i].containers[s].Name;this.res.ContainerNames.indexOf(l)===-1&&this.res.ContainerNames.push(l)}this.loading=!1})},fetchData(){this.listLoading=!0,this.listQuery.currentPage=this.page.currentPage,this.listQuery.pageSize=this.page.pageSize,T(this.listQuery).then(e=>{this.list=e.data.list,this.page=e.data.page,this.list||(this.list=[]),this.list.sort(function(a,i){return a.ServerName.localeCompare(i.ServerName)}),this.groupList=new Set(this.list.map(a=>a.ServerName)),this.groups={};for(let a=0;a<this.list.length;a++){const i=this.list[a];let s=this.groups[i.ServerName];s||(s={ServerName:i.ServerName,Members:0,StartIdx:a}),s.Members=s.Members+1,this.groups[i.ServerName]=s}this.listLoading=!1})},spanMethod(e){const{row:a,rowIndex:i,columnIndex:s}=e;if(s===1){const l=this.groups[a.ServerName];return i===l.StartIdx?{rowspan:l.Members,colspan:1}:{rowspan:0,colspan:0}}else return{rowspan:1,colspan:1}},handleSizeChange(e){this.page.pageSize=e,this.fetchData()},handleCurrentChange(e){this.page.currentPage=e,this.fetchData()},formatDate(e){return I(e)},PortsToStr(e){return B(e)},PortToStr(e){const a=(e.IP?e.IP:"")+(e.PublicPort?":"+e.PublicPort:"");return(a?a+" -> ":"")+e.PrivatePort+"/"+e.Type},FormatName(e){return E(e)},openDetail(e){this.selectRow=e,this.dialogDetailVisible=!0;const a={ContainerId:e.ContainerId};R(a).then(i=>{this.selectRow=i.data})},filterMatch(e){return!this.filterSearch.Name||e.Name.toLowerCase().includes(this.filterSearch.Name.toLowerCase())}}},j={class:"app-container"},J={class:"filter-container"},A={style:{"text-indent":"-2em"}};function G(e,a,i,s,l,u){const S=c("el-option"),N=c("el-select"),h=c("el-button"),d=c("el-table-column"),y=c("el-input"),w=c("el-table"),k=c("el-pagination"),V=c("el-dialog"),x=L("loading");return g(),f("div",j,[C("div",J,[r(N,{modelValue:l.listQuery.ServerNames,"onUpdate:modelValue":a[0]||(a[0]=t=>l.listQuery.ServerNames=t),multiple:"",filterable:"",placeholder:e.$t("主机"),clearable:"","collapse-tags":"",class:"filter-item",style:{width:"300px"}},{default:n(()=>[(g(!0),f(_,null,b(l.res.serverNames,t=>(g(),v(S,{key:t,label:t,value:t},null,8,["label","value"]))),128))]),_:1},8,["modelValue","placeholder"]),r(N,{modelValue:l.listQuery.ContainerNames,"onUpdate:modelValue":a[1]||(a[1]=t=>l.listQuery.ContainerNames=t),multiple:"",filterable:"",placeholder:e.$t("容器名称"),clearable:"","collapse-tags":"",class:"filter-item",style:{width:"300px","margin-left":"10px"}},{default:n(()=>[(g(!0),f(_,null,b(l.res.ContainerNames,t=>(g(),v(S,{key:t,label:t,value:t},null,8,["label","value"]))),128))]),_:1},8,["modelValue","placeholder"]),r(N,{modelValue:l.listQuery.state,"onUpdate:modelValue":a[2]||(a[2]=t=>l.listQuery.state=t),placeholder:"state",clearable:"",filterable:"",style:{width:"140px","margin-left":"10px","margin-right":"10px"},class:"filter-item"},{default:n(()=>[r(S,{key:"running",label:"running",value:"running"}),r(S,{key:"exited",label:"exited",value:"exited"})]),_:1},8,["modelValue"]),r(h,{class:"filter-item",type:"primary",icon:l.ElIconSearch,onClick:u.fetchData},{default:n(()=>[m(" Search ")]),_:1},8,["icon","onClick"])]),z((g(),v(w,{data:l.list.filter(t=>u.filterMatch(t)),"span-method":u.spanMethod,"element-loading-text":"Loading",border:"",fit:"","highlight-current-row":"",style:{"margin-top":"20px"}},{default:n(()=>[r(d,{align:"center",label:"ID",width:"65"},{default:n(t=>[m(o(t.$index+1),1)]),_:1}),r(d,{label:"ServerName",width:"170"},{default:n(t=>[m(o(t.row.ServerName),1)]),_:1}),r(d,{label:"Name"},{default:n(t=>[r(h,{title:t.row.Name,type:"primary",link:"",onClick:p=>u.openDetail(t.row)},{default:n(()=>[m(o(t.row.Name),1)]),_:2},1032,["title","onClick"])]),header:n(()=>[m(" Name "),r(y,{modelValue:l.filterSearch.Name,"onUpdate:modelValue":a[3]||(a[3]=t=>l.filterSearch.Name=t),size:"small",placeholder:e.$t("输入关键字过滤"),style:{width:"140px"}},null,8,["modelValue","placeholder"])]),_:1}),r(d,{label:"IMAGE",align:"center"},{default:n(t=>[C("span",null,o(t.row.Image),1)]),_:1}),r(d,{label:"Status",width:"200",align:"center"},{default:n(t=>[m(o(t.row.Status),1)]),_:1}),r(d,{label:"State",width:"80",align:"center"},{default:n(t=>[C("span",{style:M({color:t.row.State==="running"?"#03c961":"#d70404"})},o(t.row.State),5)]),_:1}),r(d,{label:"Ports",width:"270",align:"center"},{default:n(t=>[(g(!0),f(_,null,b(t.row.Ports,p=>(g(),f("div",{key:p.key},o(u.PortToStr(p)),1))),128))]),_:1}),r(d,{align:"center",prop:"created_at",label:"Created",width:"200"},{default:n(t=>[C("span",null,o(u.formatDate(t.row.Created)),1)]),_:1}),r(d,{"class-name":"status-col",label:e.$t("操作"),width:"170",align:"center"},{default:n(t=>[r(h,{type:"primary",link:"",loading:l.listLoading,disabled:t.row.State!=="running",onClick:p=>u.ContainerOperator("stop",t.row)},{default:n(()=>[m(o(e.$t("停止")),1)]),_:2},1032,["loading","disabled","onClick"]),r(h,{type:"danger",link:"",loading:l.listLoading,onClick:p=>u.ContainerOperator("remove",t.row)},{default:n(()=>[m(o(e.$t("删除")),1)]),_:2},1032,["loading","onClick"]),r(h,{type:"primary",link:"",loading:l.listLoading,onClick:p=>u.ContainerOperator("restart",t.row)},{default:n(()=>[m(o(e.$t("重启")),1)]),_:2},1032,["loading","onClick"])]),_:1},8,["label"])]),_:1},8,["data","span-method"])),[[x,l.listLoading]]),r(k,{"hide-on-single-page":!0,"current-page":l.page.currentPage,"page-sizes":[10,30,50,100,200,300,400],"page-size":l.page.pageSize,layout:"prev, pager, next, jumper, sizes, total",total:l.page.total,style:{width:"500px",margin:"10px auto 0"},onSizeChange:u.handleSizeChange,onCurrentChange:u.handleCurrentChange},null,8,["current-page","page-size","total","onSizeChange","onCurrentChange"]),r(V,{modelValue:l.dialogDetailVisible,"onUpdate:modelValue":a[4]||(a[4]=t=>l.dialogDetailVisible=t),title:e.$t("详情")},{default:n(()=>[C("pre",A,"    "+o(JSON.stringify(l.selectRow,null,2))+`
          `,1)]),_:1},8,["modelValue","title"])])}const X=D(F,[["render",G]]);export{X as default};
