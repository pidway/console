"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[7264],{47264:(e,t,a)=>{a.r(t),a.d(t,{default:()=>p});var n=a(65043),o=a(89923),l=a(56483),i=a(77403),r=a(94141),d=a(23758),s=a(64159),u=a(20554),c=a(56629),b=a(53518),h=a(70579);const p=e=>{let{open:t,enabled:a,cfg:p,selectedBucket:x,closeModalAndRefresh:g}=e;const j=(0,u.jL)(),[k,m]=(0,n.useState)(!1),[f,v]=(0,n.useState)(!1),[C,S]=(0,n.useState)("1"),[q,A]=(0,n.useState)("Ti"),[y,B]=(0,n.useState)(!1);(0,n.useEffect)((()=>{if(a&&(v(!0),p)){const e=(0,l.GT)(p.quota||0,!0,!1,!0);S(e.total.toString()),A(e.unit),B(!0)}}),[a,p]),(0,n.useEffect)((()=>{B(!f||/^\d*(?:\.\d{1,2})?$/.test(C))}),[f,C]);return(0,h.jsx)(r.A,{modalOpen:t,onClose:()=>{g()},title:"Enable Bucket Quota",titleIcon:(0,h.jsx)(o.Uh,{}),children:(0,h.jsx)("form",{noValidate:!0,autoComplete:"off",onSubmit:e=>{e.preventDefault(),!k&&y&&c.F.buckets.setBucketQuota(x,{enabled:f,amount:parseInt((0,l.q5)(C,q,!0)),quota_type:"hard"}).then((()=>{m(!1),g()})).catch((e=>{m(!1),j((0,s.Dy)((0,b.S)(e.error)))}))},children:(0,h.jsxs)(o.Hbc,{withBorders:!1,containerPadding:!1,children:[(0,h.jsx)(o.dOG,{value:"bucket_quota",id:"bucket_quota",name:"bucket_quota",checked:f,onChange:e=>{v(e.target.checked)},label:"Enabled"}),f&&(0,h.jsx)(o.cl_,{id:"quota_size",name:"quota_size",onChange:e=>{S(e.target.value),e.target.validity.valid?B(!0):B(!1)},label:"Quota",value:C,required:!0,min:"1",overlayObject:(0,h.jsx)(d.A,{id:"quota_unit",onUnitChange:e=>{A(e)},unitSelected:q,unitsList:(0,l.l9)(["Ki"]),disabled:!1}),error:y?"":"Please enter a valid quota"}),(0,h.jsxs)(o.xA9,{item:!0,xs:12,sx:i.Uz.modalButtonBar,children:[(0,h.jsx)(o.$nd,{id:"cancel",type:"button",variant:"regular",disabled:k,onClick:()=>{g()},label:"Cancel"}),(0,h.jsx)(o.$nd,{id:"save",type:"submit",variant:"callAction",disabled:k||!y,label:"Save"})]}),k&&(0,h.jsx)(o.xA9,{item:!0,xs:12,children:(0,h.jsx)(o.z21,{})})]})})})}},23758:(e,t,a)=>{a.d(t,{A:()=>u});var n=a(65043),o=a(89923),l=a(94574),i=a(33097),r=a.n(i),d=a(70579);const s=l.Ay.button((e=>{let{theme:t}=e;return{border:"1px solid ".concat(r()(t,"borderColor","#E2E2E2")),borderRadius:3,color:r()(t,"secondaryText","#5B5C5C"),backgroundColor:r()(t,"boxBackground","#FBFAFA"),fontSize:12}})),u=e=>{let{id:t,unitSelected:a,unitsList:l,disabled:i=!1,onUnitChange:r}=e;const[u,c]=n.useState(null),b=Boolean(u),h=e=>{c(null),""!==e&&r&&r(e)};return(0,d.jsxs)(n.Fragment,{children:[(0,d.jsx)(s,{id:"".concat(t,"-button"),"aria-controls":"".concat(t,"-menu"),"aria-haspopup":"true","aria-expanded":b?"true":void 0,onClick:e=>{c(e.currentTarget)},disabled:i,type:"button",children:a}),(0,d.jsx)(o.Vey,{id:"upload-main-menu",options:l,selectedOption:"",onSelect:e=>h(e),hideTriggerAction:()=>{h("")},open:b,anchorEl:u,anchorOrigin:"end"})]})}}}]);
//# sourceMappingURL=7264.f7c51a0e.chunk.js.map