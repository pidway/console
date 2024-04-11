"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[6140],{13759:(e,a,t)=>{t.r(a),t.d(a,{default:()=>o});var l=t(65043),n=t(89923),s=t(70579);const o=e=>{let{onChange:a}=e;const[t,o]=(0,l.useState)(!1),[r,u]=(0,l.useState)(""),[c,i]=(0,l.useState)(""),[d,m]=(0,l.useState)(""),[g,h]=(0,l.useState)(""),[b,p]=(0,l.useState)(""),[v,f]=(0,l.useState)(""),[x,j]=(0,l.useState)(" "),[C,S]=(0,l.useState)(""),[w,y]=(0,l.useState)("namespace"),[_,k]=(0,l.useState)(""),[E,q]=(0,l.useState)(""),[B,D]=(0,l.useState)(""),F=(0,l.useCallback)((()=>{let e="";return""!==c&&(e="".concat(e," host=").concat(c)),""!==d&&(e="".concat(e," dbname=").concat(d)),""!==b&&(e="".concat(e," user=").concat(b)),""!==v&&(e="".concat(e," password=").concat(v)),""!==g&&(e="".concat(e," port=").concat(g))," "!==x&&(e="".concat(e," sslmode=").concat(x)),e="".concat(e," "),e.trim()}),[c,d,b,v,g,x]);return(0,l.useEffect)((()=>{if(""!==r){a([{key:"connection_string",value:r},{key:"table",value:C},{key:"format",value:w},{key:"queue_dir",value:_},{key:"queue_limit",value:E},{key:"comment",value:B}])}}),[r,C,w,_,E,B,a]),(0,l.useEffect)((()=>{const e=F();u(e)}),[b,d,v,g,x,c,u,F]),(0,l.useEffect)((()=>{if(t){const e=F();return void u(e)}const e=((e,a)=>{let t=[];for(const s of a){const a=e.indexOf(s+"=");-1!==a&&t.push(a)}t.sort(((e,a)=>e-a));let l=new Map,n=new Array(t.length);for(let s=0;s<t.length;s++){const a=s+1;a<t.length?n[s]=e.slice(t[s],t[a]):n[s]=e.slice(t[s])}for(let s of n){if(void 0===s)continue;const e=s.slice(0,s.indexOf("=")),a=s.slice(s.indexOf("=")+1).trim();l.set(e,a)}return l})(r,["host","port","dbname","user","password","sslmode"]);i(e.get("host")?e.get("host")+"":""),h(e.get("port")?e.get("port")+"":""),m(e.get("dbname")?e.get("dbname")+"":""),p(e.get("user")?e.get("user")+"":""),f(e.get("password")?e.get("password")+"":""),j(e.get("sslmode")?e.get("sslmode")+"":" ")}),[t]),(0,s.jsxs)(n.Hbc,{containerPadding:!1,withBorders:!1,children:[(0,s.jsx)(n.dOG,{label:"Manually Configure String",checked:t,id:"manualString",name:"manualString",onChange:e=>{o(e.target.checked)},value:"manualString"}),t?(0,s.jsx)(l.Fragment,{children:(0,s.jsx)(n.cl_,{id:"connection-string",name:"connection_string",label:"Connection String",value:r,onChange:e=>{u(e.target.value)}})}):(0,s.jsxs)(l.Fragment,{children:[(0,s.jsx)(n.xA9,{item:!0,xs:12,children:(0,s.jsxs)(n.azJ,{withBorders:!0,useBackground:!0,sx:{overflowY:"auto",height:170,marginBottom:12},children:[(0,s.jsx)(n.cl_,{id:"host",name:"host",label:"",placeholder:"Enter Host",value:c,onChange:e=>{i(e.target.value)}}),(0,s.jsx)(n.cl_,{id:"db-name",name:"db-name",label:"",placeholder:"Enter DB Name",value:d,onChange:e=>{m(e.target.value)}}),(0,s.jsx)(n.cl_,{id:"port",name:"port",label:"",placeholder:"Enter Port",value:g,onChange:e=>{h(e.target.value)}}),(0,s.jsx)(n.l6P,{value:x,label:"",id:"sslmode",name:"sslmode",onChange:e=>{e&&j(e+"")},options:[{label:"Enter SSL Mode",value:" "},{label:"Require",value:"require"},{label:"Disable",value:"disable"},{label:"Verify CA",value:"verify-ca"},{label:"Verify Full",value:"verify-full"}]}),(0,s.jsx)(n.cl_,{id:"user",name:"user",label:"",placeholder:"Enter User",value:b,onChange:e=>{p(e.target.value)}}),(0,s.jsx)(n.cl_,{id:"password",name:"password",label:"",type:"password",placeholder:"Enter Password",value:v,onChange:e=>{f(e.target.value)}})]})}),(0,s.jsx)(n.EmB,{label:"Connection String",multiLine:!0,children:r})]}),(0,s.jsx)(n.cl_,{id:"table",name:"table",label:"Table",placeholder:"Enter Table Name",value:C,tooltip:"DB table name to store/update events, table is auto-created",onChange:e=>{S(e.target.value)}}),(0,s.jsx)(n.z6M,{currentValue:w,id:"format",name:"format",label:"Format",onChange:e=>{y(e.target.value)},tooltip:"'namespace' reflects current bucket/object list and 'access' reflects a journal of object operations, defaults to 'namespace'",selectorOptions:[{label:"Namespace",value:"namespace"},{label:"Access",value:"access"}]}),(0,s.jsx)(n.cl_,{id:"queue-dir",name:"queue_dir",label:"Queue Dir",placeholder:"Enter Queue Directory",value:_,tooltip:"Staging directory for undelivered messages e.g. '/home/events'",onChange:e=>{k(e.target.value)}}),(0,s.jsx)(n.cl_,{id:"queue-limit",name:"queue_limit",label:"Queue Limit",placeholder:"Enter Queue Limit",type:"number",value:E,tooltip:"Maximum limit for undelivered messages, defaults to '10000'",onChange:e=>{q(e.target.value)}}),(0,s.jsx)(n.hFj,{id:"comment",name:"comment",label:"Comment",placeholder:"Enter custom notes if any",value:B,onChange:e=>{D(e.target.value)}})]})}}}]);
//# sourceMappingURL=6140.37801ce8.chunk.js.map