"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[5851],{85851:(e,t,n)=>{n.r(t),n.d(t,{default:()=>h});var o=n(65043),l=n(89923),s=n(56483),r=n(64159),c=n(20554),a=n(58661),i=n(56629),u=n(53518),p=n(70579);const h=e=>{let{selectedGroups:t,deleteOpen:n,closeDeleteModalAndRefresh:h}=e;const f=(0,c.jL)(),[d,g]=(0,o.useState)(!1);if(!t)return null;const j=t.map((e=>(0,p.jsx)("div",{children:(0,p.jsx)("b",{children:e})},e)));return(0,p.jsx)(a.A,{title:"Delete Group".concat(t.length>1?"s":""),confirmText:"Delete",isOpen:n,titleIcon:(0,p.jsx)(l.xWY,{}),isLoading:d,onConfirm:()=>{for(let e of t)g(!0),i.F.group.removeGroup((0,s.nf)(e)).then((e=>{h(!0)})).catch((async e=>{const t=await e.json();f((0,r.C9)((0,u.S)(t))),h(!1)})).finally((()=>g(!1)))},onClose:()=>h(!1),confirmationContent:(0,p.jsxs)(o.Fragment,{children:["Are you sure you want to delete the following"," ",1===t.length?"":t.length," group",t.length>1?"s?":"?",j]})})}}}]);
//# sourceMappingURL=5851.9d7a7887.chunk.js.map