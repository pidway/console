"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[5809],{5809:(e,l,t)=>{t.r(l),t.d(l,{default:()=>_});var s=t(65043),a=t(33097),n=t.n(a),c=t(89923),o=t(99161),d=t(77938),r=t(64159),i=t(20554),u=t(25448),p=t(72237),x=t(70579);const T=(0,p.A)(s.lazy((()=>t.e(1303).then(t.bind(t,1303))))),h=(0,p.A)(s.lazy((()=>t.e(5079).then(t.bind(t,55079))))),_=e=>{let{bucketName:l}=e;const t=(0,i.jL)(),[a,p]=(0,s.useState)(null),[_,v]=(0,s.useState)(!1),[b,j]=(0,s.useState)([]),[O,S]=(0,s.useState)(["",""]),[A,C]=(0,s.useState)(!1),[g,G]=(0,u.A)((e=>{if(e&&null!=(null===e||void 0===e?void 0:e.details)){var l,t;if(e.details.tags)return p(null===e||void 0===e||null===(l=e.details)||void 0===l?void 0:l.tags),void j(Object.keys(null===e||void 0===e||null===(t=e.details)||void 0===t?void 0:t.tags));p([]),j([])}}),(e=>{t((0,r.C9)(e))})),f=()=>{G("GET","/api/v1/buckets/".concat(l))};return(0,s.useEffect)((()=>{f()}),[l]),(0,x.jsxs)(c.azJ,{children:[g?(0,x.jsx)(c.aHM,{style:{width:16,height:16}}):null,(0,x.jsx)(d.R,{scopes:[o.OV.S3_GET_BUCKET_TAGGING,o.OV.S3_GET_ACTIONS],resource:l,children:(0,x.jsx)(c.azJ,{sx:{display:"flex",flexFlow:"column",marginTop:5},children:(0,x.jsxs)(c.azJ,{sx:{display:"flex",gap:8,flexWrap:"wrap"},children:[b&&b.map(((e,t)=>{const s=n()(a,"".concat(e),"");return""!==s?(0,x.jsx)(d.R,{scopes:[o.OV.S3_PUT_BUCKET_TAGGING,o.OV.S3_PUT_ACTIONS],resource:l,matchAll:!0,errorProps:{deleteIcon:null,onDelete:null},children:(0,x.jsx)(c.vwO,{label:"".concat(e," : ").concat(s),id:"tag-".concat(e,"-").concat(s),onDelete:()=>{((e,l)=>{S([e,l]),C(!0)})(e,s)}})},"chip-".concat(t)):null})),(0,x.jsx)(d.R,{scopes:[o.OV.S3_PUT_BUCKET_TAGGING,o.OV.S3_PUT_ACTIONS],resource:l,errorProps:{disabled:!0,onClick:null},children:(0,x.jsx)(c.vwO,{label:"Add tag",icon:(0,x.jsx)(c.REV,{}),id:"create-tag",variant:"outlined",onClick:()=>{v(!0)},sx:{cursor:"pointer",maxWidth:90}})})]})})}),_&&(0,x.jsx)(T,{modalOpen:_,currentTags:a,bucketName:l,onCloseAndUpdate:e=>{v(!1),e&&f()}}),A&&(0,x.jsx)(h,{deleteOpen:A,currentTags:a,bucketName:l,onCloseAndUpdate:e=>{C(!1),e&&f()},selectedTag:O})]})}}}]);
//# sourceMappingURL=5809.064e83cc.chunk.js.map