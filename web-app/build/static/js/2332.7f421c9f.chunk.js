"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[2332],{32332:(e,t,n)=>{n.r(t),n.d(t,{default:()=>j,emptyContent:()=>y});var a=n(65043),s=n(89923),l=n(73216),o=n(99161),r=n(64159),c=n(20554),i=n(77403),p=n(38375),x=n(48793),m=n(92452),d=n(56629),u=n(53518),h=n(70579);const y='{\n    "bytes": ""\n}',j=()=>{const e=(0,c.jL)(),t=(0,l.Zp)(),[n,j]=(0,a.useState)(!1),[b,k]=(0,a.useState)(""),[f,C]=(0,a.useState)(y),S=""!==b.trim()&&-1===b.indexOf(" ");return(0,a.useEffect)((()=>{e((0,r.ph)("import_key"))}),[]),(0,h.jsx)(a.Fragment,{children:(0,h.jsxs)(s.xA9,{item:!0,xs:12,children:[(0,h.jsx)(x.A,{label:(0,h.jsx)(s.EGL,{onClick:()=>t(o.zZ.KMS_KEYS),label:"Keys"}),actions:(0,h.jsx)(m.A,{})}),(0,h.jsx)(s.Mxu,{children:(0,h.jsx)(s.Hbc,{title:"Import Key",icon:(0,h.jsx)(s.No_,{}),helpBox:(0,h.jsx)(p.A,{helpText:"Encryption Key",contents:["Import a cryptographic key in the Key Management Service server connected to MINIO."]}),children:(0,h.jsxs)("form",{noValidate:!0,autoComplete:"off",onSubmit:n=>{(n=>{j(!0),n.preventDefault();let a=JSON.parse(f);d.F.kms.kmsImportKey(b,a).then((e=>{t("".concat(o.zZ.KMS_KEYS))})).catch((async t=>{const n=await t.json();e((0,r.C9)((0,u.S)(n)))})).finally((()=>j(!1)))})(n)},children:[(0,h.jsx)(s.cl_,{id:"key-name",name:"key-name",label:"Key Name",autoFocus:!0,value:b,error:(e=>-1!==e.indexOf(" ")?"Key name cannot contain spaces":"")(b),onChange:e=>{k(e.target.value)}}),(0,h.jsx)(s.BYM,{label:"Set key Content",value:f,onChange:e=>{C(e)},editorHeight:"350px"}),(0,h.jsxs)(s.xA9,{item:!0,xs:12,sx:i.Uz.modalButtonBar,children:[(0,h.jsx)(s.$nd,{id:"clear",type:"button",variant:"regular",onClick:()=>{k(""),C("")},label:"Clear"}),(0,h.jsx)(s.$nd,{id:"import-key",type:"submit",variant:"callAction",color:"primary",disabled:n||!S,label:"Import"})]})]})})})]})})}},38375:(e,t,n)=>{n.d(t,{A:()=>o});var a=n(65043),s=n(89923),l=n(70579);const o=e=>{let{helpText:t,contents:n}=e;return(0,l.jsx)(s.lVp,{iconComponent:(0,l.jsx)(s.nag,{}),title:t,help:(0,l.jsx)(a.Fragment,{children:n.map((e=>(0,l.jsx)(s.azJ,{sx:{paddingBottom:"20px"},children:e})))})})}}}]);
//# sourceMappingURL=2332.7f421c9f.chunk.js.map