"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[9714],{79714:(e,t,a)=>{a.r(t),a.d(t,{default:()=>h});var l=a(65043),n=a(89923),s=a(94141),i=a(96512),c=a(77403),r=a(64159),o=a(20554),d=a(56629),p=a(53518),u=a(70579);const h=e=>{let{closeModalAndRefresh:t,open:a,bucketName:h,ruleID:g}=e;const x=(0,o.jL)(),[m,f]=(0,l.useState)(!0),[j,S]=(0,l.useState)(!1),[b,v]=(0,l.useState)("1"),[k,y]=(0,l.useState)(""),[C,D]=(0,l.useState)(""),[E,R]=(0,l.useState)(!1),[I,w]=(0,l.useState)(!1),[_,A]=(0,l.useState)(""),[B,O]=(0,l.useState)(""),[M,N]=(0,l.useState)(""),[P,T]=(0,l.useState)(!1),[G,F]=(0,l.useState)(!1),[V,q]=(0,l.useState)(!1);return(0,l.useEffect)((()=>{m&&d.F.buckets.getBucketReplicationRule(h,g).then((e=>{var t;v(e.data.priority?e.data.priority.toString():"");const a=e.data.prefix||"",l=e.data.tags||"";D(a),A(l),O(l),y((null===(t=e.data.destination)||void 0===t?void 0:t.bucket)||""),R(e.data.delete_marker_replication||!1),N(e.data.storageClass||""),T(!!e.data.existingObjects),F(!!e.data.deletes_replication),q("Enabled"===e.data.status),w(!!e.data.metadata_replication),f(!1)})).catch((e=>{x((0,r.Dy)((0,p.S)(e.error))),f(!1)}))}),[m,x,h,g]),(0,l.useEffect)((()=>{if(j){const e={arn:k,ruleState:V,prefix:C,tags:B,replicateDeleteMarkers:E,replicateDeletes:G,replicateExistingObjects:P,replicateMetadata:I,priority:parseInt(b),storageClass:M};d.F.buckets.updateMultiBucketReplication(h,g,e).then((()=>{S(!1),t(!0)})).catch((e=>{x((0,r.Dy)((0,p.S)(e.error))),S(!1)}))}}),[j,h,g,k,C,B,E,b,G,P,V,I,M,t,x]),(0,u.jsx)(s.A,{modalOpen:a,onClose:()=>{t(!1)},title:"Edit Bucket Replication",titleIcon:(0,u.jsx)(n.WBh,{}),children:(0,u.jsx)("form",{noValidate:!0,autoComplete:"off",onSubmit:e=>{e.preventDefault(),S(!0)},children:(0,u.jsxs)(n.Hbc,{containerPadding:!1,withBorders:!1,children:[(0,u.jsx)(n.dOG,{checked:V,id:"ruleState",name:"ruleState",label:"Rule State",onChange:e=>{q(e.target.checked)}}),(0,u.jsx)(n.EmB,{label:"Destination",sx:{width:"100%"},children:k}),(0,u.jsx)(n.cl_,{id:"priority",name:"priority",onChange:e=>{e.target.validity.valid&&v(e.target.value)},label:"Priority",value:b,pattern:"[0-9]*"}),(0,u.jsx)(n.cl_,{id:"storageClass",name:"storageClass",onChange:e=>{N(e.target.value)},placeholder:"STANDARD_IA,REDUCED_REDUNDANCY etc",label:"Storage Class",value:M}),(0,u.jsxs)("fieldset",{className:"inputItem",children:[(0,u.jsx)("legend",{children:"Object Filters"}),(0,u.jsx)(n.cl_,{id:"prefix",name:"prefix",onChange:e=>{D(e.target.value)},placeholder:"prefix",label:"Prefix",value:C}),(0,u.jsx)(i.A,{name:"tags",label:"Tags",elements:_,onChange:e=>{O(e)},keyPlaceholder:"Tag Key",valuePlaceholder:"Tag Value",withBorder:!0})]}),(0,u.jsxs)("fieldset",{className:"inputItem",children:[(0,u.jsx)("legend",{children:"Replication Options"}),(0,u.jsx)(n.dOG,{checked:P,id:"repExisting",name:"repExisting",label:"Existing Objects",onChange:e=>{T(e.target.checked)},description:"Replicate existing objects"}),(0,u.jsx)(n.dOG,{checked:I,id:"metadatataSync",name:"metadatataSync",label:"Metadata Sync",onChange:e=>{w(e.target.checked)},description:"Metadata Sync"}),(0,u.jsx)(n.dOG,{checked:E,id:"deleteMarker",name:"deleteMarker",label:"Delete Marker",onChange:e=>{R(e.target.checked)},description:"Replicate soft deletes"}),(0,u.jsx)(n.dOG,{checked:G,id:"repDelete",name:"repDelete",label:"Deletes",onChange:e=>{F(e.target.checked)},description:"Replicate versioned deletes"})]}),(0,u.jsxs)(n.xA9,{item:!0,xs:12,sx:c.Uz.modalButtonBar,children:[(0,u.jsx)(n.$nd,{id:"cancel-edit-replication",type:"button",variant:"regular",disabled:m||j,onClick:()=>{t(!1)},label:"Cancel"}),(0,u.jsx)(n.$nd,{id:"save-replication",type:"submit",variant:"callAction",disabled:m||j,label:"Save"})]})]})})})}},96512:(e,t,a)=>{a.d(t,{A:()=>d});var l=a(65043),n=a(33097),s=a.n(n),i=a(93950),c=a.n(i),r=a(89923),o=a(70579);const d=e=>{let{elements:t,name:a,label:n,tooltip:i="",keyPlaceholder:d="",valuePlaceholder:p="",onChange:u,withBorder:h=!1}=e;const[g,x]=(0,l.useState)([""]),[m,f]=(0,l.useState)([""]),j=(0,l.createRef)();(0,l.useEffect)((()=>{if(1===g.length&&""===g[0]&&1===m.length&&""===m[0]&&t&&""!==t){const e=t.split("&");let a=[],l=[];e.forEach((e=>{const t=e.split("=");2===t.length&&(a.push(t[0]),l.push(t[1]))})),a.push(""),l.push(""),x(a),f(l)}}),[g,m,t]),(0,l.useEffect)((()=>{const e=j.current;e&&g.length>1&&e.scrollIntoView(!1)}),[g]);const S=(0,l.useRef)(!0);(0,l.useLayoutEffect)((()=>{S.current?S.current=!1:k()}),[g,m]);const b=e=>{e.persist();let t=[...g];const a=s()(e.target,"dataset.index","0");t[parseInt(a)]=e.target.value,x(t)},v=e=>{e.persist();let t=[...m];const a=s()(e.target,"dataset.index","0");t[parseInt(a)]=e.target.value,f(t)},k=c()((()=>{let e="";g.forEach(((t,a)=>{if(g[a]&&m[a]){let l="".concat(t,"=").concat(m[a]);0!==a&&(l="&".concat(l)),e="".concat(e).concat(l)}})),u(e)}),500),y=m.map(((e,t)=>(0,o.jsxs)(r.xA9,{item:!0,xs:12,className:"lineInputBoxes inputItem",children:[(0,o.jsx)(r.cl_,{id:"".concat(a,"-key-").concat(t.toString()),label:"",name:"".concat(a,"-").concat(t.toString()),value:g[t],onChange:b,index:t,placeholder:d}),(0,o.jsx)("span",{className:"queryDiv",children:":"}),(0,o.jsx)(r.cl_,{id:"".concat(a,"-value-").concat(t.toString()),label:"",name:"".concat(a,"-").concat(t.toString()),value:m[t],onChange:v,index:t,placeholder:p,overlayIcon:t===m.length-1?(0,o.jsx)(r.REV,{}):null,overlayAction:()=>{(()=>{if(""!==g[g.length-1].trim()&&""!==m[m.length-1].trim()){const e=[...g],t=[...m];e.push(""),t.push(""),x(e),f(t)}})()}})]},"query-pair-".concat(a,"-").concat(t.toString()))));return(0,o.jsx)(l.Fragment,{children:(0,o.jsxs)(r.xA9,{item:!0,xs:12,sx:{"& .lineInputBoxes":{display:"flex"},"& .queryDiv":{alignSelf:"center",margin:"-15px 4px 0",fontWeight:600}},className:"inputItem",children:[(0,o.jsxs)(r.l1Y,{children:[n,""!==i&&(0,o.jsx)(r.azJ,{sx:{marginLeft:5,display:"flex",alignItems:"center","& .min-icon":{width:13}},children:(0,o.jsx)(r.m_M,{tooltip:i,placement:"top",children:(0,o.jsx)(r.NTw,{style:{width:13,height:13}})})})]}),(0,o.jsxs)(r.azJ,{withBorders:h,sx:{padding:15,height:150,overflowY:"auto",position:"relative",marginTop:15},children:[y,(0,o.jsx)("div",{ref:j})]})]})})}}}]);
//# sourceMappingURL=9714.48a29c42.chunk.js.map