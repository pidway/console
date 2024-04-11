"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[2166],{72271:(e,t,a)=>{a.d(t,{A:()=>i});a(65043);var n=a(89923),s=a(70579);const i=e=>{let{iconComponent:t,entity:a}=e;return(0,s.jsx)(n.xA9,{container:!0,children:(0,s.jsx)(n.xA9,{item:!0,xs:12,children:(0,s.jsx)(n.lVp,{title:"".concat(a," not available"),iconComponent:t,help:(0,s.jsxs)(n.azJ,{sx:{fontSize:"14px",["@media (max-width: ".concat(n.nmC.sm,"px)")]:{display:"flex",flexFlow:"column"}},children:[(0,s.jsx)("span",{children:"This feature is not available for a single-disk setup.\xa0"}),(0,s.jsxs)("span",{children:["Please deploy a server in"," ",(0,s.jsx)("a",{href:"https://min.io/docs/minio/linux/operations/install-deploy-manage/deploy-minio-multi-node-multi-drive.html?ref=con",target:"_blank",rel:"noopener",children:"Distributed Mode"})," ","to use this feature."]})]})})})})}},2166:(e,t,a)=>{a.r(t),a.d(t,{default:()=>g});var n=a(65043),s=a(89923),i=a(73216),l=a(22166),o=a(56483),c=a(64159),r=a(20554),d=a(14558),p=a(94141),x=a(72271),m=a(49501),h=a(28481),u=a(48793),f=a(92452),j=a(70579);const b=e=>{let{volumeVal:t,pathVal:a}=e;return(0,j.jsx)(s.azJ,{className:"code-block-container",children:(0,j.jsxs)(s.azJ,{className:"example-code-block",children:[(0,j.jsxs)(s.azJ,{sx:{display:"flex",marginBottom:"5px",flexFlow:"row",["@media (max-width: ".concat(s.nmC.sm,"px)")]:{flexFlow:"column"}},children:[(0,j.jsx)("label",{children:"Volume/bucket Name :"})," ",(0,j.jsx)("code",{children:t})]}),(0,j.jsxs)(s.azJ,{sx:{display:"flex",flexFlow:"row",["@media (max-width: ".concat(s.nmC.sm,"px)")]:{flexFlow:"column"}},children:[(0,j.jsx)("label",{children:"Path : "}),(0,j.jsx)("code",{children:a})]})]})})},g=()=>{const e=(0,r.jL)(),t=(0,i.Zp)(),a=(0,l.d4)(c.Rq),[g,w]=(0,n.useState)(""),[y,v]=(0,n.useState)(""),[k,z]=(0,n.useState)(!0),[C,T]=(0,n.useState)(""),[S,J]=(0,n.useState)(""),[F,V]=(0,n.useState)(!1),[A,_]=(0,n.useState)(""),[N,I]=(0,n.useState)(""),B=(0,d.vf)();(0,n.useEffect)((()=>{let e,t;e=g.trim().length>0,e?"/"===g.slice(0,1)&&(e=!1,_("Volume/Bucket name cannot start with /")):_("This field is required"),t=y.trim().length>0,y?"/"===y.slice(0,1)&&(t=!1,I("Path cannot start with /")):I("This field is required");const a=e&&t;e&&_(""),t&&I(""),V(a)}),[g,y]);const E=async()=>{const t=(0,o.nf)(y),a=(0,o.nf)(g);let n=document.baseURI.replace(window.location.origin,"");(async e=>await fetch(e,{method:"GET"}))("".concat(n,"/api/v1/admin/inspect?volume=").concat(a,"&file=").concat(t,"&encrypt=").concat(k)).then((async t=>{if(!t.ok){const a=await t.json();e((0,c.C9)({errorMessage:a.message,detailedError:a.code}))}const a=await t.blob(),n=t.headers.get("content-disposition").split('"')[1],s=(0,o.UM)(n)||"";(0,o.OT)(a,n),J(n),T(s)})).catch((t=>{e((0,c.C9)(t))}))},L=()=>{w(""),v(""),z(!0)};return(0,n.useEffect)((()=>{e((0,c.ph)("inspect"))}),[]),(0,j.jsxs)(n.Fragment,{children:[(0,j.jsx)(u.A,{label:"Inspect",actions:(0,j.jsx)(f.A,{})}),(0,j.jsxs)(s.Mxu,{children:[!B&&(0,j.jsx)(h.A,{compactMode:!0}),a?(0,j.jsx)(s.Hbc,{helpBox:(0,j.jsx)(s.lVp,{title:"Learn more about the Inspect feature",iconComponent:(0,j.jsx)(s.nTF,{}),help:(0,j.jsxs)(n.Fragment,{children:[(0,j.jsx)(s.azJ,{sx:{marginTop:"16px",fontWeight:600,fontStyle:"italic",fontSize:"14px"},children:"Examples:"}),(0,j.jsxs)(s.azJ,{sx:{display:"flex",flexFlow:"column",fontSize:"14px",flex:"2","& .step-row":{fontSize:"14px",display:"flex",marginTop:"15px",marginBottom:"15px","&.step-text":{fontWeight:400},"&:before":{content:"' '",height:"7px",width:"7px",backgroundColor:"#2781B0",marginRight:"10px",marginTop:"7px",flexShrink:0}},"& .code-block-container":{flex:"1",marginTop:"15px",marginLeft:"35px","& input":{color:"#737373"}},"& .example-code-block label":{display:"inline-block",width:160,fontWeight:600,fontSize:14,["@media (max-width: ".concat(s.nmC.sm,"px)")]:{width:"100%"}},"& code":{width:100,paddingLeft:"10px",fontFamily:"monospace",paddingRight:"10px",paddingTop:"3px",paddingBottom:"3px",borderRadius:"2px",border:"1px solid #eaeaea",fontSize:"10px",fontWeight:500,["@media (max-width: ".concat(s.nmC.sm,"px)")]:{width:"100%"}},"& .spacer":{marginBottom:"5px"}},children:[(0,j.jsxs)(s.azJ,{children:[(0,j.jsx)(s.azJ,{className:"step-row",children:(0,j.jsx)("div",{className:"step-text",children:"To Download 'xl.meta' for a specific object from all the drives in a zip file:"})}),(0,j.jsx)(b,{pathVal:"test*/xl.meta",volumeVal:"test-bucket"})]}),(0,j.jsxs)(s.azJ,{children:[(0,j.jsx)(s.azJ,{className:"step-row",children:(0,j.jsx)("div",{className:"step-text",children:"To Download all constituent parts for a specific object, and optionally encrypt the downloaded zip:"})}),(0,j.jsx)(b,{pathVal:"test*/xl.meta",volumeVal:"test*/*/part.*"})]}),(0,j.jsxs)(s.azJ,{children:[(0,j.jsx)(s.azJ,{className:"step-row",children:(0,j.jsxs)("div",{className:"step-text",children:["To Download recursively all objects at a prefix.",(0,j.jsx)("br",{}),"NOTE: This can be an expensive operation use it with caution."]})}),(0,j.jsx)(b,{pathVal:"test*/xl.meta",volumeVal:"test/**"})]})]}),(0,j.jsxs)(s.azJ,{sx:{marginTop:"30px",marginLeft:"15px",fontSize:"14px"},children:["You can learn more at our"," ",(0,j.jsx)("a",{href:"https://github.com/minio/minio/tree/master/docs/debugging?ref=con",target:"_blank",rel:"noopener",children:"documentation"}),"."]})]})}),children:(0,j.jsxs)("form",{noValidate:!0,autoComplete:"off",onSubmit:e=>{e.preventDefault(),B?E():t("/support/register")},children:[(0,j.jsx)(s.cl_,{id:"inspect_volume",name:"inspect_volume",onChange:e=>{w(e.target.value)},label:"Volume or Bucket Name",value:g,error:A,required:!0,placeholder:"test-bucket",disabled:!B}),(0,j.jsx)(s.cl_,{id:"inspect_path",name:"inspect_path",error:N,onChange:e=>{v(e.target.value)},label:"File or Path to inspect",value:y,required:!0,placeholder:"test*/xl.meta",disabled:!B}),(0,j.jsx)(s.dOG,{label:"Encrypt",indicatorLabels:["True","False"],checked:k,value:"true",id:"inspect_encrypt",name:"inspect_encrypt",onChange:()=>{z(!k)},disabled:!B}),(0,j.jsxs)(s.azJ,{sx:{display:"flex",alignItems:"center",justifyContent:"flex-end",marginTop:"55px"},children:[(0,j.jsx)(s.$nd,{id:"inspect-clear-button",style:{marginRight:"15px"},type:"button",variant:"regular","data-test-id":"inspect-clear-button",onClick:L,label:"Clear",disabled:!B}),(0,j.jsx)(s.$nd,{id:"inspect-start",type:"submit",variant:B?"callAction":"regular","data-test-id":"inspect-submit-button",disabled:!F||!B,label:"Inspect"})]})]})}):(0,j.jsx)(x.A,{iconComponent:(0,j.jsx)(s.nTF,{}),entity:"Inspect"}),C?(0,j.jsx)(p.A,{modalOpen:!0,title:"Inspect Decryption Key",onClose:()=>{(0,o.Yj)(S),T(""),L()},titleIcon:(0,j.jsx)(s.aJN,{}),children:(0,j.jsxs)(n.Fragment,{children:[(0,j.jsxs)(s.azJ,{children:["This will be displayed only once. It cannot be recovered.",(0,j.jsx)("br",{}),"Use secure medium to share this key."]}),(0,j.jsx)("form",{noValidate:!0,onSubmit:()=>!1,children:(0,j.jsx)(m.A,{value:C})})]})}):null]})]})}},49501:(e,t,a)=>{a.d(t,{A:()=>l});var n=a(65043),s=a(89923),i=a(70579);const l=e=>{let{value:t}=e;const[a,l]=(0,n.useState)(!1);return(0,i.jsxs)(s.azJ,{sx:{display:"flex",alignItems:"center",flexFlow:"row",["@media (max-width: ".concat(s.nmC.sm,"px)")]:{flexFlow:"column"}},children:[(0,i.jsx)(s.cl_,{id:"inspect-dec-key",name:"inspect-dec-key",placeholder:"",label:"",type:a?"text":"password",onChange:()=>{},value:t,overlayIcon:(0,i.jsx)(s.TdU,{}),readOnly:!0,overlayAction:()=>navigator.clipboard.writeText(t)}),(0,i.jsx)(s.$nd,{id:"show-hide-key",style:{marginLeft:"10px"},variant:"callAction",onClick:()=>l(!a),label:"Show/Hide"})]})}}}]);
//# sourceMappingURL=2166.1e6c2b39.chunk.js.map