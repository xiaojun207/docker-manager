(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-7a707321"],{"0ccb":function(t,e,n){var r=n("50c4"),a=n("1148"),i=n("1d80"),o=Math.ceil,l=function(t){return function(e,n,l){var s,u,c=String(i(e)),f=c.length,d=void 0===l?" ":String(l),g=r(n);return g<=f||""==d?c:(s=g-f,u=a.call(d,o(s/d.length)),u.length>s&&(u=u.slice(0,s)),t?c+u:u+c)}};t.exports={start:l(!1),end:l(!0)}},1148:function(t,e,n){"use strict";var r=n("a691"),a=n("1d80");t.exports="".repeat||function(t){var e=String(a(this)),n="",i=r(t);if(i<0||i==1/0)throw RangeError("Wrong number of repetitions");for(;i>0;(i>>>=1)&&(e+=e))1&i&&(n+=e);return n}},"25f0":function(t,e,n){"use strict";var r=n("6eeb"),a=n("825a"),i=n("d039"),o=n("ad6d"),l="toString",s=RegExp.prototype,u=s[l],c=i((function(){return"/a/b"!=u.call({source:"a",flags:"b"})})),f=u.name!=l;(c||f)&&r(RegExp.prototype,l,(function(){var t=a(this),e=String(t.source),n=t.flags,r=String(void 0===n&&t instanceof RegExp&&!("flags"in s)?o.call(t):n);return"/"+e+"/"+r}),{unsafe:!0})},"408a":function(t,e,n){var r=n("c6b6");t.exports=function(t){if("number"!=typeof t&&"Number"!=r(t))throw TypeError("Incorrect invocation");return+t}},"43db":function(t,e,n){"use strict";n.r(e);var r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"app-container"},[n("div",{staticClass:"filter-container"},[n("el-select",{staticClass:"filter-item",staticStyle:{width:"300px"},attrs:{multiple:"",filterable:"",placeholder:t.$t("服务器"),clearable:"","collapse-tags":""},model:{value:t.listQuery.serverNames,callback:function(e){t.$set(t.listQuery,"serverNames",e)},expression:"listQuery.serverNames"}},t._l(t.res.serverNames,(function(t){return n("el-option",{key:t,attrs:{label:t,value:t}})})),1),n("el-select",{staticClass:"filter-item",staticStyle:{width:"300px","margin-left":"10px"},attrs:{multiple:"",filterable:"",placeholder:t.$t("容器名称"),clearable:"","collapse-tags":""},model:{value:t.listQuery.ContainerNames,callback:function(e){t.$set(t.listQuery,"ContainerNames",e)},expression:"listQuery.ContainerNames"}},t._l(t.res.ContainerNames,(function(t){return n("el-option",{key:t,attrs:{label:t,value:t}})})),1),n("el-select",{staticClass:"filter-item",staticStyle:{width:"140px","margin-left":"10px","margin-right":"10px"},attrs:{placeholder:t.$t("是否开启日志"),clearable:"",filterable:""},model:{value:t.listQuery.Follow,callback:function(e){t.$set(t.listQuery,"Follow",e)},expression:"listQuery.Follow"}},[n("el-option",{key:"true",attrs:{label:t.$t("是"),value:"true"}}),n("el-option",{key:"false",attrs:{label:t.$t("否"),value:"false"}})],1),n("el-button",{staticClass:"filter-item",attrs:{type:"primary",icon:"el-icon-search"},on:{click:t.fetchData}},[t._v(" Search ")])],1),n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.listLoading,expression:"listLoading"}],staticStyle:{"margin-top":"15px"},attrs:{data:t.list,"element-loading-text":"Loading",border:"",fit:"","highlight-current-row":""}},[n("el-table-column",{attrs:{align:"center",label:"ID",width:"95"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.$index+1)+" ")]}}])}),n("el-table-column",{attrs:{label:"Name",width:"210"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-button",{attrs:{type:"text"},on:{click:function(n){return t.openDetail(e.row)}}},[t._v(t._s(e.row.Name))])]}}])}),n("el-table-column",{attrs:{label:"ServerName",width:"170",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("span",[t._v(t._s(e.row.ServerName))])]}}])}),n("el-table-column",{attrs:{label:"CpuStats",width:"150",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(t.formatCpu(e.row.cpu_stats,e.row.precpu_stats))+" ")]}}])}),n("el-table-column",{attrs:{label:"MemoryStats","min-width":"180",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(t.formatMemory(e.row.memory_stats))+" ")]}}])}),n("el-table-column",{attrs:{label:"MEM %","min-width":"150",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(t.formatMemory2(e.row.memory_stats))+" ")]}}])}),n("el-table-column",{attrs:{label:"Networks",width:"210",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(t.formatNet(e.row.networks))+" ")]}}])}),n("el-table-column",{attrs:{label:t.$t("实时日志"),width:"100",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[e.row.Follow?n("span",{staticStyle:{color:"#03c961"}},[t._v(t._s(t.$t("已开启")))]):t._e(),e.row.Follow?t._e():n("span",{staticStyle:{color:"#d70404"}},[t._v(t._s(t.$t("未开启")))])]}}])}),n("el-table-column",{attrs:{label:"UpdateDate",width:"210",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(t.formatDate(e.row.UpdateDate))+" ")]}}])}),n("el-table-column",{attrs:{"class-name":"status-col",label:t.$t("操作"),width:"220",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-button",{attrs:{loading:t.listLoading,disabled:e.row.Follow,type:"text"},on:{click:function(n){return t.startLog(e.row)}}},[t._v(t._s(t.$t("开启日志")))]),n("el-button",{attrs:{loading:t.listLoading,disabled:!e.row.Follow,type:"text"},on:{click:function(n){return t.closeLog(e.row)}}},[t._v(t._s(t.$t("关闭日志")))])]}}])})],1),n("el-dialog",{attrs:{visible:t.dialogDetailVisible,title:t.$t("详情")},on:{"update:visible":function(e){t.dialogDetailVisible=e}}},[n("pre",[t._v(t._s(JSON.stringify(t.selectRow,null,2))+"\n      ")])])],1)},a=[],i=n("ed08"),o=n("b775");function l(t){return Object(o["a"])({url:"/mgr/stats",method:"get",params:t})}var s=n("4dd0"),u=n("62ed"),c=n("b1cf"),f={filters:{statusFilter:function(t){var e={published:"success",draft:"gray",deleted:"danger"};return e[t]}},data:function(){return{list:null,listLoading:!0,dialogDetailVisible:!1,selectRow:{},res:{serverNames:[],ContainerNames:[],containerInfos:[]},listQuery:{serverNames:[],ContainerNames:[],Follow:void 0}}},created:function(){this.fetchServerNames(),this.fetchContainerInfoData(),this.fetchData()},methods:{fetchServerNames:function(){var t=this;this.listLoading=!0,Object(s["e"])().then((function(e){t.res.serverNames=e.data,t.listLoading=!1}))},fetchContainerInfoData:function(){var t=this;this.loading=!0,Object(s["b"])().then((function(e){var n=e.data;for(var r in t.res.ContainerNames=[],n)for(var a in n[r].containers){var i=n[r].containers[a].Name;-1===t.res.ContainerNames.indexOf(i)&&t.res.ContainerNames.push(i)}t.loading=!1}))},fetchData:function(){var t=this;this.listLoading=!0,l(this.listQuery).then((function(e){t.list=e.data,t.listLoading=!1}))},startLog:function(t){var e=this,n={containerId:t.ContainerId};this.loading=!0,Object(u["b"])(n).then((function(t){e.loading=!1,e.$message(e.$t("命令已下发"))})).catch((function(t){e.loading=!1,e.$message(t.msg)}))},closeLog:function(t){var e=this,n={containerId:t.ContainerId};this.loading=!0,Object(u["a"])(n).then((function(t){e.loading=!1,e.$message(e.$t("命令已下发"))})).catch((function(t){e.loading=!1}))},openDetail:function(t){this.selectRow=t,this.dialogDetailVisible=!0},formatDate:function(t){return Object(i["a"])(t)},formatMemory:function(t){return Object(c["e"])(t)},formatMemory2:function(t){return Object(c["f"])(t)},formatCpu:function(t,e){return Object(c["d"])(t,e)},formatNet:function(t){return Object(c["g"])(t)}}},d=f,g=n("2877"),p=Object(g["a"])(d,r,a,!1,null,null,null);e["default"]=p.exports},"4d63":function(t,e,n){var r=n("83ab"),a=n("da84"),i=n("94ca"),o=n("7156"),l=n("9bf2").f,s=n("241c").f,u=n("44e7"),c=n("ad6d"),f=n("9f7f"),d=n("6eeb"),g=n("d039"),p=n("69f3").set,b=n("2626"),v=n("b622"),h=v("match"),m=a.RegExp,y=m.prototype,_=/a/g,w=/a/g,S=new m(_)!==_,x=f.UNSUPPORTED_Y,N=r&&i("RegExp",!S||x||g((function(){return w[h]=!1,m(_)!=_||m(w)==w||"/a/i"!=m(_,"i")})));if(N){var $=function(t,e){var n,r=this instanceof $,a=u(t),i=void 0===e;if(!r&&a&&t.constructor===$&&i)return t;S?a&&!i&&(t=t.source):t instanceof $&&(i&&(e=c.call(t)),t=t.source),x&&(n=!!e&&e.indexOf("y")>-1,n&&(e=e.replace(/y/g,"")));var l=o(S?new m(t,e):m(t,e),r?this:y,$);return x&&n&&p(l,{sticky:n}),l},k=function(t){t in $||l($,t,{configurable:!0,get:function(){return m[t]},set:function(e){m[t]=e}})},O=s(m),C=0;while(O.length>C)k(O[C++]);y.constructor=$,$.prototype=y,d(a,"RegExp",$)}b("RegExp")},"4d90":function(t,e,n){"use strict";var r=n("23e7"),a=n("0ccb").start,i=n("9a0c");r({target:"String",proto:!0,forced:i},{padStart:function(t){return a(this,t,arguments.length>1?arguments[1]:void 0)}})},"4dd0":function(t,e,n){"use strict";n.d(e,"d",(function(){return a})),n.d(e,"e",(function(){return i})),n.d(e,"c",(function(){return o})),n.d(e,"b",(function(){return l})),n.d(e,"f",(function(){return s})),n.d(e,"a",(function(){return u}));var r=n("b775");function a(t){return Object(r["a"])({url:"/mgr/servers",method:"get",params:t})}function i(t){return Object(r["a"])({url:"/mgr/serverNames",method:"get",params:t})}function o(t){return Object(r["a"])({url:"/mgr/containers",method:"get",params:t})}function l(t){return Object(r["a"])({url:"/mgr/containerInfos",method:"get",params:t})}function s(t){return Object(r["a"])({url:"/mgr/publish",method:"post",data:t})}function u(t,e){return Object(r["a"])({url:"/mgr/container/"+t,method:"post",data:e})}},5319:function(t,e,n){"use strict";var r=n("d784"),a=n("825a"),i=n("7b0b"),o=n("50c4"),l=n("a691"),s=n("1d80"),u=n("8aa5"),c=n("14c3"),f=Math.max,d=Math.min,g=Math.floor,p=/\$([$&'`]|\d\d?|<[^>]*>)/g,b=/\$([$&'`]|\d\d?)/g,v=function(t){return void 0===t?t:String(t)};r("replace",2,(function(t,e,n,r){var h=r.REGEXP_REPLACE_SUBSTITUTES_UNDEFINED_CAPTURE,m=r.REPLACE_KEEPS_$0,y=h?"$":"$0";return[function(n,r){var a=s(this),i=void 0==n?void 0:n[t];return void 0!==i?i.call(n,a,r):e.call(String(a),n,r)},function(t,r){if(!h&&m||"string"===typeof r&&-1===r.indexOf(y)){var i=n(e,t,this,r);if(i.done)return i.value}var s=a(t),g=String(this),p="function"===typeof r;p||(r=String(r));var b=s.global;if(b){var w=s.unicode;s.lastIndex=0}var S=[];while(1){var x=c(s,g);if(null===x)break;if(S.push(x),!b)break;var N=String(x[0]);""===N&&(s.lastIndex=u(g,o(s.lastIndex),w))}for(var $="",k=0,O=0;O<S.length;O++){x=S[O];for(var C=String(x[0]),E=f(d(l(x.index),g.length),0),M=[],D=1;D<x.length;D++)M.push(v(x[D]));var j=x.groups;if(p){var F=[C].concat(M,E,g);void 0!==j&&F.push(j);var I=String(r.apply(void 0,F))}else I=_(C,g,E,M,j,r);E>=k&&($+=g.slice(k,E)+I,k=E+C.length)}return $+g.slice(k)}];function _(t,n,r,a,o,l){var s=r+t.length,u=a.length,c=b;return void 0!==o&&(o=i(o),c=p),e.call(l,c,(function(e,i){var l;switch(i.charAt(0)){case"$":return"$";case"&":return t;case"`":return n.slice(0,r);case"'":return n.slice(s);case"<":l=o[i.slice(1,-1)];break;default:var c=+i;if(0===c)return e;if(c>u){var f=g(c/10);return 0===f?e:f<=u?void 0===a[f-1]?i.charAt(1):a[f-1]+i.charAt(1):e}l=a[c-1]}return void 0===l?"":l}))}}))},"53ca":function(t,e,n){"use strict";n.d(e,"a",(function(){return r}));n("a4d3"),n("e01a"),n("d3b7"),n("d28b"),n("3ca3"),n("ddb0");function r(t){return r="function"===typeof Symbol&&"symbol"===typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"===typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t},r(t)}},"62ed":function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"a",(function(){return i}));var r=n("b775");function a(t){return console.log("getLogStart:",t),Object(r["a"])({url:"/mgr/log/start",method:"post",data:t})}function i(t){return Object(r["a"])({url:"/mgr/log/close",method:"post",data:t})}},7156:function(t,e,n){var r=n("861d"),a=n("d2bb");t.exports=function(t,e,n){var i,o;return a&&"function"==typeof(i=e.constructor)&&i!==n&&r(o=i.prototype)&&o!==n.prototype&&a(t,o),t}},"9a0c":function(t,e,n){var r=n("342f");t.exports=/Version\/10\.\d+(\.\d+)?( Mobile\/\w+)? Safari\//.test(r)},a15b:function(t,e,n){"use strict";var r=n("23e7"),a=n("44ad"),i=n("fc6a"),o=n("a640"),l=[].join,s=a!=Object,u=o("join",",");r({target:"Array",proto:!0,forced:s||!u},{join:function(t){return l.call(i(this),void 0===t?",":t)}})},b1cf:function(t,e,n){"use strict";n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return a})),n.d(e,"b",(function(){return i})),n.d(e,"h",(function(){return o})),n.d(e,"e",(function(){return l})),n.d(e,"f",(function(){return s})),n.d(e,"d",(function(){return u})),n.d(e,"g",(function(){return c}));n("b680");function r(t){if(!t||0===t.length)return"";var e="";for(var n in t){var r=t[n],a=(r.IP?r.IP:"")+(r.PublicPort?":"+r.PublicPort:"");e+=", "+(a?a+"->":"")+r.PrivatePort+"/"+r.Type}return e.substr(1,e.length)}function a(t){return t[0].substr(1,t[0].length-1)}function i(t){if(!t||0===t.length)return"";var e=[];for(var n in t){var r=t[n];e.push(r.Source+" : "+r.Destination)}return e}function o(t){if(!t)return"";var e="";return e=t<1024?t+"B":t<1048576?(t/1024).toFixed(3)+"KB":t<1073741824?(t/1048576).toFixed(3)+"MB":(t/1073741824).toFixed(3)+"GB",e}function l(t){return t&&t.usage?o(t.usage)+" / "+o(t.limit):""}function s(t){return t&&t.usage?(100*t.usage/t.limit).toFixed(2)+"%":""}function u(t,e){if(!t||!t.cpu_usage||!e||!e.cpu_usage||!t.system_cpu_usage||!e.system_cpu_usage)return"0%";var n=t.cpu_usage.total_usage-e.cpu_usage.total_usage,r=t.system_cpu_usage-e.system_cpu_usage,a=t.online_cpus,i=n/r*a*100;return i.toFixed(2)+"%"}function c(t){if(!t||!t.eth0)return"";var e=t.eth0;return o(e.rx_bytes)+" / "+o(e.tx_bytes)}},b680:function(t,e,n){"use strict";var r=n("23e7"),a=n("a691"),i=n("408a"),o=n("1148"),l=n("d039"),s=1..toFixed,u=Math.floor,c=function(t,e,n){return 0===e?n:e%2===1?c(t,e-1,n*t):c(t*t,e/2,n)},f=function(t){var e=0,n=t;while(n>=4096)e+=12,n/=4096;while(n>=2)e+=1,n/=2;return e},d=s&&("0.000"!==8e-5.toFixed(3)||"1"!==.9.toFixed(0)||"1.25"!==1.255.toFixed(2)||"1000000000000000128"!==(0xde0b6b3a7640080).toFixed(0))||!l((function(){s.call({})}));r({target:"Number",proto:!0,forced:d},{toFixed:function(t){var e,n,r,l,s=i(this),d=a(t),g=[0,0,0,0,0,0],p="",b="0",v=function(t,e){var n=-1,r=e;while(++n<6)r+=t*g[n],g[n]=r%1e7,r=u(r/1e7)},h=function(t){var e=6,n=0;while(--e>=0)n+=g[e],g[e]=u(n/t),n=n%t*1e7},m=function(){var t=6,e="";while(--t>=0)if(""!==e||0===t||0!==g[t]){var n=String(g[t]);e=""===e?n:e+o.call("0",7-n.length)+n}return e};if(d<0||d>20)throw RangeError("Incorrect fraction digits");if(s!=s)return"NaN";if(s<=-1e21||s>=1e21)return String(s);if(s<0&&(p="-",s=-s),s>1e-21)if(e=f(s*c(2,69,1))-69,n=e<0?s*c(2,-e,1):s/c(2,e,1),n*=4503599627370496,e=52-e,e>0){v(0,n),r=d;while(r>=7)v(1e7,0),r-=7;v(c(10,r,1),0),r=e-1;while(r>=23)h(1<<23),r-=23;h(1<<r),v(1,1),h(2),b=m()}else v(0,n),v(1<<-e,0),b=m()+o.call("0",d);return d>0?(l=b.length,b=p+(l<=d?"0."+o.call("0",d-l)+b:b.slice(0,l-d)+"."+b.slice(l-d))):b=p+b,b}})},d28b:function(t,e,n){var r=n("746f");r("iterator")},e01a:function(t,e,n){"use strict";var r=n("23e7"),a=n("83ab"),i=n("da84"),o=n("5135"),l=n("861d"),s=n("9bf2").f,u=n("e893"),c=i.Symbol;if(a&&"function"==typeof c&&(!("description"in c.prototype)||void 0!==c().description)){var f={},d=function(){var t=arguments.length<1||void 0===arguments[0]?void 0:String(arguments[0]),e=this instanceof d?new c(t):void 0===t?c():c(t);return""===t&&(f[e]=!0),e};u(d,c);var g=d.prototype=c.prototype;g.constructor=d;var p=g.toString,b="Symbol(test)"==String(c("test")),v=/^Symbol\((.*)\)[^)]+$/;s(g,"description",{configurable:!0,get:function(){var t=l(this)?this.valueOf():this,e=p.call(t);if(o(f,t))return"";var n=b?e.slice(7,-1):e.replace(v,"$1");return""===n?void 0:n}}),r({global:!0,forced:!0},{Symbol:d})}},ed08:function(t,e,n){"use strict";n.d(e,"a",(function(){return r}));n("53ca"),n("ac1f"),n("5319"),n("4d63"),n("25f0"),n("d3b7"),n("4d90"),n("a15b"),n("d81d"),n("b64b"),n("1276"),n("159b");function r(t){if(!t)return"";t=t>0xee09da7916?t:1e3*t;var e=new Date(t),n=e.getFullYear()+"-",r=(e.getMonth()+1<10?"0"+(e.getMonth()+1):e.getMonth()+1)+"-",a=e.getDate()<10?"0"+e.getDate():e.getDate(),i=(e.getHours()<10?"0"+e.getHours():e.getHours())+":",o=(e.getMinutes()<10?"0"+e.getMinutes():e.getMinutes())+":",l=(e.getSeconds()<10?"0"+e.getSeconds():e.getSeconds())+".",s="";return s=e.getMilliseconds()<10?"00"+e.getMilliseconds():e.getMilliseconds()<100?"0"+e.getMilliseconds():e.getMilliseconds(),n+r+a+" "+i+o+l+s}}}]);