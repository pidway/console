"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[9361],{19361:(e,r,t)=>{t.r(r),t.d(r,{default:()=>m});var a=t(65043),o=t(94574),n=t(73216),i=t(20649),s=t(64710),l=t(89923),c=t(14558),d=t(33097),p=t.n(d),g=t(70579);const h=o.Ay.div((e=>{let{theme:r}=e;return{"& .errorDescription":{fontStyle:"italic",transition:"all .2s ease-in-out",padding:"0 15px",marginTop:5,overflow:"auto"},"& .errorLabel":{color:p()(r,"fontColor","#000"),fontSize:18,fontWeight:"bold",marginLeft:5},"& .simpleError":{marginTop:5,padding:"2px 5px",fontSize:16,color:p()(r,"fontColor","#000")},"& .messageIcon":{color:p()(r,"signalColors.danger","#C72C48"),display:"flex","& svg":{width:32,height:32}},"& .errorTitle":{display:"flex",alignItems:"center",borderBottom:15}}})),m=()=>{const e=(0,n.Zp)(),[r,t]=(0,a.useState)(""),[o,d]=(0,a.useState)(""),[p,m]=(0,a.useState)(!0);return(0,a.useEffect)((()=>{if(p){const r=window.location.search,a=new URLSearchParams(r),o=a.get("code"),n=a.get("state"),s=a.get("error"),l=a.get("errorDescription");s||l?(t(s||""),d(l||""),m(!1)):i.A.invoke("POST","/api/v1/login/oauth2/auth",{code:o,state:n}).then((()=>{let r="/";localStorage.getItem("redirect-path")&&""!==localStorage.getItem("redirect-path")&&(r="".concat(localStorage.getItem("redirect-path")),localStorage.setItem("redirect-path","")),n&&localStorage.setItem("auth-state",n),m(!1),e(r)})).catch((e=>{t(e.errorMessage),d(e.detailedError),m(!1)}))}}),[p,e]),""!==r||""!==o?(0,g.jsx)(a.Fragment,{children:(0,g.jsx)(l.ndn,{logoProps:{applicationName:"console",subVariant:(0,c.vP)()},form:(0,g.jsxs)(h,{children:[(0,g.jsxs)("div",{className:"errorTitle",children:[(0,g.jsx)("span",{className:"messageIcon",children:(0,g.jsx)(l.cJw,{})}),(0,g.jsx)("span",{className:"errorLabel",children:"Error from IDP"})]}),(0,g.jsx)("div",{className:"simpleError",children:r}),(0,g.jsx)(l.azJ,{className:"errorDescription",children:o}),(0,g.jsx)(l.$nd,{id:"back-to-login",onClick:()=>{window.location.href="".concat(s.p,"login")},type:"submit",variant:"callAction",fullWidth:!0,children:"Back to Login"})]}),promoHeader:(0,g.jsx)("span",{style:{fontSize:28},children:"High-Performance Object Store"}),promoInfo:(0,g.jsxs)("span",{style:{fontSize:14,lineHeight:1},children:["MinIO is a cloud-native object store built to run on any infrastructure - public, private or edge clouds. Primary use cases include data lakes, databases, AI/ML, SaaS applications and fast backup & recovery. MinIO is dual licensed under GNU AGPL v3 and commercial license. To learn more, visit"," ",(0,g.jsx)("a",{href:"https://min.io/?ref=con",target:"_blank",rel:"noopener",children:"www.min.io"}),"."]}),backgroundAnimation:!1})}):null}}}]);
//# sourceMappingURL=9361.3fc638a6.chunk.js.map