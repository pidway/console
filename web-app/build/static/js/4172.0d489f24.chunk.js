"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[4172],{4172:(e,t,n)=>{n.r(t),n.d(t,{default:()=>p});var s=n(65043),c=n(33097),l=n.n(c),r=n(25448),i=n(58661),o=n(89923),u=n(64159),a=n(20554),f=n(70579);const p=e=>{let{closeDeleteModalAndRefresh:t,deleteOpen:n,selectedBucket:c,bucketEvent:p}=e;const d=(0,a.jL)(),[x,v]=(0,r.A)((()=>t(!0)),(e=>d((0,u.C9)(e))));if(!c)return null;return(0,f.jsx)(i.A,{title:"Delete Event",confirmText:"Delete",isOpen:n,titleIcon:(0,f.jsx)(o.xWY,{}),isLoading:x,onConfirm:()=>{if(null===p)return;const e=l()(p,"events",[]),t=l()(p,"prefix",""),n=l()(p,"suffix",""),s=e.reduce(((e,t)=>e.includes(t)?e:[...e,t]),[]);v("DELETE","/api/v1/buckets/".concat(c,"/events/").concat(p.arn),{events:s,prefix:t,suffix:n})},onClose:()=>t(!1),confirmationContent:(0,f.jsx)(s.Fragment,{children:"Are you sure you want to delete this event?"})})}}}]);
//# sourceMappingURL=4172.0d489f24.chunk.js.map