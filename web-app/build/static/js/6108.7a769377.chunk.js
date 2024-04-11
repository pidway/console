"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[6108],{66108:(e,t,a)=>{a.r(t),a.d(t,{default:()=>x});var l=a(65043),s=a(89923),n=a(64159),i=a(20554),c=a(56629),o=a(85330),r=a(94141),u=a(77403),d=a(70579);const x=e=>{let{open:t,selectedBucket:a,closeModalAndRefresh:x}=e;const b=(0,i.jL)(),[h,m]=(0,l.useState)(!1),[p,f]=(0,l.useState)(""),[j,v]=(0,l.useState)(""),[g,A]=(0,l.useState)(""),[S,k]=(0,l.useState)([]),[C,E]=(0,l.useState)([]),R=(0,l.useCallback)((()=>{m(!0),c.F.admin.arnList().then((e=>{null!==e.data.arns&&E(e.data.arns),m(!1)})).catch((e=>{m(!1),b((0,n.Dy)(e))}))}),[b]);(0,l.useEffect)((()=>{R()}),[R]);const w=[{label:"PUT - Object Uploaded",value:o.Wj.Put},{label:"GET - Object accessed",value:o.Wj.Get},{label:"DELETE - Object Deleted",value:o.Wj.Delete},{label:"REPLICA - Object Replicated",value:o.Wj.Replica},{label:"ILM - Object Transitioned",value:o.Wj.Ilm},{label:"SCANNER - Object has too many versions / Prefixes has too many sub-folders",value:o.Wj.Scanner}],y=null===C||void 0===C?void 0:C.map((e=>({label:e,value:e})));return(0,d.jsx)(r.A,{modalOpen:t,onClose:()=>{x()},title:"Subscribe To Bucket Events",titleIcon:(0,d.jsx)(s.VDx,{}),children:(0,d.jsx)("form",{noValidate:!0,autoComplete:"off",onSubmit:e=>{e.preventDefault(),h||(m(!0),c.F.buckets.createBucketEvent(a,{configuration:{arn:g,events:S,prefix:p,suffix:j},ignoreExisting:!0}).then((()=>{m(!1),x()})).catch((e=>{m(!1),b((0,n.Dy)(e))})))},children:(0,d.jsxs)(s.xA9,{container:!0,children:[(0,d.jsxs)(s.xA9,{item:!0,xs:12,sx:u.a_.formScrollable,children:[(0,d.jsx)(s.xA9,{item:!0,xs:12,sx:{...u.h$.formFieldRow,"& div div .MuiOutlinedInput-root":{padding:0}},children:(0,d.jsx)(s.jT8,{onChange:e=>{A(e)},id:"select-access-policy",name:"select-access-policy",label:"ARN",value:g,options:y||[],helpTip:(0,d.jsx)(l.Fragment,{children:(0,d.jsx)("a",{target:"blank",href:"https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html",children:"Amazon Resource Name"})})})}),(0,d.jsx)(s.xA9,{item:!0,xs:12,sx:u.h$.formFieldRow,children:(0,d.jsx)(s.cl_,{id:"prefix-input",name:"prefix-input",label:"Prefix",value:p,onChange:e=>{f(e.target.value)}})}),(0,d.jsx)(s.xA9,{item:!0,xs:12,sx:u.h$.formFieldRow,children:(0,d.jsx)(s.cl_,{id:"suffix-input",name:"suffix-input",label:"Suffix",value:j,onChange:e=>{v(e.target.value)}})}),(0,d.jsx)(s.xA9,{item:!0,xs:12,sx:u.h$.formFieldRow,children:(0,d.jsx)(s.bQt,{columns:[{label:"Event",elementKey:"label"}],idField:"value",records:w,onSelect:e=>{const t=e.target,a=t.value,l=t.checked;let s=[...S];l?s.push(a):s=s.filter((e=>e!==a)),k(s)},selectedItems:S,noBackground:!0,customPaperHeight:"260px"})})]}),(0,d.jsxs)(s.xA9,{item:!0,xs:12,sx:u.Uz.modalButtonBar,children:[(0,d.jsx)(s.$nd,{id:"cancel-add-event",type:"button",variant:"regular",disabled:h,onClick:()=>{x()},label:"Cancel"}),(0,d.jsx)(s.$nd,{id:"save-event",type:"submit",variant:"callAction",disabled:h||""===g||0===S.length,label:"Save"})]})]})})})}}}]);
//# sourceMappingURL=6108.7a769377.chunk.js.map