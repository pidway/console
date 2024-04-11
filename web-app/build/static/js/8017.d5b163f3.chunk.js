"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[8017],{38017:(e,t,o)=>{o.r(t),o.d(t,{default:()=>A});var n=o(65043),a=o(89923),i=o(22166),r=o(73216),l=o(56629),s=o(53518),c=o(94574),p=o(33097),d=o.n(p),u=o(67433),m=o(77403),g=o(99161),h=o(64159),f=o(20554),y=o(39161),b=o(64681),S=o(58661),v=o(70579);const T=e=>{let{onConfirm:t,onClose:o,serviceName:i,status:r}=e;return(0,v.jsx)(S.A,{title:"Delete Endpoint",confirmText:"Delete",isOpen:!0,titleIcon:(0,v.jsx)(a.$rg,{}),isLoading:!1,onConfirm:t,onClose:o,confirmationContent:(0,v.jsxs)(n.Fragment,{children:["Are you sure you want to delete the event destination ?",(0,v.jsx)("br",{}),(0,v.jsx)("b",{children:i})," which is ",(0,v.jsx)("b",{children:r})]})})};var x=o(6681);const _=c.Ay.div((e=>{let{theme:t}=e;return{display:"flex",alignItems:"center","& svg":{width:16,marginRight:5,fill:d()(t,"signalColors.good","#4CCB92")},"& svg.offline":{fill:d()(t,"signalColors.danger","#C51B3F")}}})),A=()=>{const e=(0,f.jL)(),t=(0,r.Zp)(),o=(0,i.d4)((e=>e.destination.loading)),[c,p]=(0,n.useState)([]),[d,S]=(0,n.useState)(""),[A,k]=(0,n.useState)(!1),[E,j]=(0,n.useState)();(0,n.useEffect)((()=>{if(o){(()=>{l.F.admin.notificationEndpointList().then((t=>{let o=[];t.data.notification_endpoints&&(o=t.data.notification_endpoints),p((0,u.Es)(o)),e((0,y.$Y)(!1))})).catch((t=>{e((0,h.C9)((0,s.S)(t.error))),e((0,y.$Y)(!1))}))})()}}),[o,e]),(0,n.useEffect)((()=>{e((0,y.$Y)(!0))}),[e]);const C=[{type:"delete",onClick:e=>{j(e),k(!0)}}],L=c.filter((e=>""===d||e.service_name.indexOf(d)>=0));return(0,v.jsx)(n.Fragment,{children:(0,v.jsxs)(a.Mxu,{children:[(0,v.jsxs)(a.xA9,{item:!0,xs:12,sx:m._0.actionsTray,children:[(0,v.jsx)(b.A,{placeholder:"Search target",onChange:S,value:d,sx:{maxWidth:380}}),(0,v.jsxs)(a.azJ,{sx:{display:"flex",alignItems:"center",justifyContent:"flex-end",gap:5},children:[(0,v.jsx)(x.A,{tooltip:"Refresh List",children:(0,v.jsx)(a.$nd,{id:"reload-event-destinations",label:"Refresh",variant:"regular",icon:(0,v.jsx)(a.fNY,{}),onClick:()=>{e((0,y.$Y)(!0))}})}),(0,v.jsx)(x.A,{tooltip:"Add Event Destination",children:(0,v.jsx)(a.$nd,{id:"add-notification-target",label:"Add Event Destination",variant:"callAction",icon:(0,v.jsx)(a.REV,{}),onClick:()=>{t(g.zZ.EVENT_DESTINATIONS_ADD)}})})]})]}),o&&(0,v.jsx)(a.z21,{}),!o&&(0,v.jsxs)(n.Fragment,{children:[c.length>0&&(0,v.jsxs)(n.Fragment,{children:[(0,v.jsx)(a.azJ,{sx:{width:"100%"},children:(0,v.jsx)(a.bQt,{itemActions:C,columns:[{label:"Status",elementKey:"status",renderFunction:e=>(0,v.jsxs)(_,{children:[(0,v.jsx)(a.GQ2,{className:"Offline"===e?"offline":""}),e]}),width:150},{label:"Service",elementKey:"service_name"}],isLoading:o,records:L,entityName:"Event Destinations",idField:"service_name",customPaperHeight:"400px"})}),(0,v.jsx)(a.xA9,{item:!0,xs:12,sx:{marginTop:15},children:(0,v.jsx)(a.lVp,{title:"Event Destinations",iconComponent:(0,v.jsx)(a.PI5,{}),help:(0,v.jsxs)(n.Fragment,{children:["MinIO bucket notifications allow administrators to send notifications to supported external services on certain object or bucket events. MinIO supports bucket and object-level S3 events similar to the Amazon S3 Event Notifications.",(0,v.jsx)("br",{}),(0,v.jsx)("br",{}),"You can learn more at our"," ",(0,v.jsx)("a",{href:"https://min.io/docs/minio/linux/administration/monitoring/bucket-notifications.html?ref=con",target:"_blank",rel:"noopener",children:"documentation"}),"."]})})})]}),0===c.length&&(0,v.jsx)(a.xA9,{container:!0,sx:{justifyContent:"center",alignContent:"center",alignItems:"center"},children:(0,v.jsx)(a.xA9,{item:!0,xs:8,children:(0,v.jsx)(a.lVp,{title:"Event Destinations",iconComponent:(0,v.jsx)(a.PI5,{}),help:(0,v.jsxs)(n.Fragment,{children:["MinIO bucket notifications allow administrators to send notifications to supported external services on certain object or bucket events. MinIO supports bucket and object-level S3 events similar to the Amazon S3 Event Notifications.",(0,v.jsx)("br",{}),(0,v.jsx)("br",{}),"To get started,"," ",(0,v.jsx)(a.t53,{onClick:()=>{t(g.zZ.EVENT_DESTINATIONS_ADD)},children:"Add an Event Destination"}),"."]})})})})]}),A?(0,v.jsx)(T,{onConfirm:()=>{(t=>{if(null!==t&&void 0!==t&&t.name){const o=(0,u.h4)(t.name);let n=":".concat(t.account_id);o?l.F.configs.resetConfig("".concat(o).concat(n)).then((()=>{e((0,h.YR)(!0)),j(null),k(!1),e((0,y.$Y)(!0))})).catch((t=>{k(!1),e((0,h.C9)((0,s.S)(t.error)))})):(j(null),k(!1),console.log("Unable to find Config key for ".concat(t.name)))}})(E)},status:"".concat(null===E||void 0===E?void 0:E.status),serviceName:"".concat(null===E||void 0===E?void 0:E.service_name),onClose:()=>{k(!1)}}):null]})})}},67433:(e,t,o)=>{o.d(t,{AU:()=>a,D3:()=>g,Es:()=>m,P4:()=>n,Xm:()=>b,bo:()=>f,fx:()=>S,h4:()=>T});const n="notify_postgres",a="notify_mysql",i="notify_kafka",r="notify_amqp",l="notify_mqtt",s="notify_redis",c="notify_nats",p="notify_elasticsearch",d="notify_webhook",u="notify_nsq",m=e=>e.map((e=>({service_name:"".concat(e.service,":").concat(e.account_id),name:e.service,account_id:e.account_id,status:e.status})));class g{}g.DB="database",g.Queue="queue",g.Func="functions";const h=()=>"".concat(document.baseURI),f=[{actionTrigger:n,targetTitle:"PostgreSQL",logo:"".concat(h(),"postgres-logo.svg"),category:g.DB},{actionTrigger:i,targetTitle:"Kafka",logo:"".concat(h(),"kafka-logo.svg"),category:g.Queue},{actionTrigger:r,targetTitle:"AMQP",logo:"".concat(h(),"amqp-logo.svg"),category:g.Queue},{actionTrigger:l,targetTitle:"MQTT",logo:"".concat(h(),"mqtt-logo.svg"),category:g.Queue},{actionTrigger:s,targetTitle:"Redis",logo:"".concat(h(),"redis-logo.svg"),category:g.Queue},{actionTrigger:c,targetTitle:"NATS",logo:"".concat(h(),"nats-logo.svg"),category:g.Queue},{actionTrigger:a,targetTitle:"Mysql",logo:"".concat(h(),"mysql-logo.svg"),category:g.DB},{actionTrigger:p,targetTitle:"Elastic Search",logo:"".concat(h(),"elasticsearch-logo.svg"),category:g.DB},{actionTrigger:d,targetTitle:"Webhook",logo:"".concat(h(),"webhooks-logo.svg"),category:g.Func},{actionTrigger:u,targetTitle:"NSQ",logo:"".concat(h(),"nsq-logo.svg"),category:g.Queue}],y=[{name:"queue_dir",label:"Queue Directory",required:!1,tooltip:"Staging directory for undelivered messages e.g. '/home/events'",type:"string",placeholder:"Enter Queue Directory"},{name:"queue_limit",label:"Queue Limit",required:!1,tooltip:"Maximum limit for undelivered messages, defaults to '10000'",type:"number",placeholder:"Enter Queue Limit"},{name:"comment",label:"Comment",required:!1,type:"comment",placeholder:"Enter custom notes if any"}],b=e=>e.filter((e=>""!==e.value)),S={[i]:[{name:"brokers",label:"Brokers",required:!0,tooltip:"Comma separated list of Kafka broker addresses",type:"string",placeholder:"Enter Brokers"},{name:"topic",label:"Topic",tooltip:"Kafka topic used for bucket notifications",type:"string",placeholder:"Enter Topic"},{name:"sasl_username",label:"SASL Username",tooltip:"Username for SASL/PLAIN or SASL/SCRAM authentication",type:"string",placeholder:"Enter SASL Username"},{name:"sasl_password",label:"SASL Password",tooltip:"Password for SASL/PLAIN or SASL/SCRAM authentication",type:"string",placeholder:"Enter SASL Password"},{name:"sasl_mechanism",label:"SASL Mechanism",tooltip:"SASL authentication mechanism, default 'PLAIN'",type:"string"},{name:"tls_client_auth",label:"TLS Client Auth",tooltip:"Client Auth determines the Kafka server's policy for TLS client authorization",type:"string",placeholder:"Enter TLS Client Auth"},{name:"sasl",label:"SASL",tooltip:"Set to 'on' to enable SASL authentication",type:"on|off"},{name:"tls",label:"TLS",tooltip:"Set to 'on' to enable TLS",type:"on|off"},{name:"tls_skip_verify",label:"TLS skip verify",tooltip:'Trust server TLS without verification, defaults to "on" (verify)',type:"on|off"},{name:"client_tls_cert",label:"client TLS cert",tooltip:"Path to client certificate for mTLS authorization",type:"path",placeholder:"Enter TLS Client Cert"},{name:"client_tls_key",label:"client TLS key",tooltip:"Path to client key for mTLS authorization",type:"path",placeholder:"Enter TLS Client Key"},{name:"version",label:"Version",tooltip:"Specify the version of the Kafka cluster e.g '2.2.0'",type:"string",placeholder:"Enter Kafka Version"},...y],[r]:[{name:"url",required:!0,label:"URL",tooltip:"AMQP server endpoint e.g. `amqp://myuser:mypassword@localhost:5672`",type:"url"},{name:"exchange",label:"Exchange",tooltip:"Name of the AMQP exchange",type:"string",placeholder:"Enter Exchange"},{name:"exchange_type",label:"Exchange Type",tooltip:"AMQP exchange type",type:"string",placeholder:"Enter Exchange Type"},{name:"routing_key",label:"Routing Key",tooltip:"Routing key for publishing",type:"string",placeholder:"Enter Routing Key"},{name:"mandatory",label:"Mandatory",tooltip:"Quietly ignore undelivered messages when set to 'off', default is 'on'",type:"on|off"},{name:"durable",label:"Durable",tooltip:"Persist queue across broker restarts when set to 'on', default is 'off'",type:"on|off"},{name:"no_wait",label:"No Wait",tooltip:"Non-blocking message delivery when set to 'on', default is 'off'",type:"on|off"},{name:"internal",label:"Internal",tooltip:"Set to 'on' for exchange to be not used directly by publishers, but only when bound to other exchanges",type:"on|off"},{name:"auto_deleted",label:"Auto Deleted",tooltip:"Auto delete queue when set to 'on', when there are no consumers",type:"on|off"},{name:"delivery_mode",label:"Delivery Mode",tooltip:"Set to '1' for non-persistent or '2' for persistent queue",type:"number",placeholder:"Enter Delivery Mode"},...y],[s]:[{name:"address",required:!0,label:"Address",tooltip:"Redis server's address e.g. `localhost:6379`",type:"address",placeholder:"Enter Address"},{name:"key",required:!0,label:"Key",tooltip:"Redis key to store/update events, key is auto-created",type:"string",placeholder:"Enter Key"},{name:"password",label:"Password",tooltip:"Redis server password",type:"string",placeholder:"Enter Password"},...y],[l]:[{name:"broker",required:!0,label:"Broker",tooltip:"MQTT server endpoint e.g. `tcp://localhost:1883`",type:"uri",placeholder:"Enter Brokers"},{name:"topic",required:!0,label:"Topic",tooltip:"Name of the MQTT topic to publish",type:"string",placeholder:"Enter Topic"},{name:"username",label:"Username",tooltip:"MQTT username",type:"string",placeholder:"Enter Username"},{name:"password",label:"Password",tooltip:"MQTT password",type:"string",placeholder:"Enter Password"},{name:"qos",label:"QOS",tooltip:"Set the quality of service priority, defaults to '0'",type:"number",placeholder:"Enter QOS"},{name:"keep_alive_interval",label:"Keep Alive Interval",tooltip:"Keep-alive interval for MQTT connections in s,m,h,d",type:"duration",placeholder:"Enter Keep Alive Interval"},{name:"reconnect_interval",label:"Reconnect Interval",tooltip:"Reconnect interval for MQTT connections in s,m,h,d",type:"duration",placeholder:"Enter Reconnect Interval"},...y],[c]:[{name:"address",required:!0,label:"Address",tooltip:"NATS server address e.g. '0.0.0.0:4222'",type:"address",placeholder:"Enter Address"},{name:"subject",required:!0,label:"Subject",tooltip:"NATS subscription subject",type:"string",placeholder:"Enter NATS Subject"},{name:"username",label:"Username",tooltip:"NATS username",type:"string",placeholder:"Enter NATS Username"},{name:"password",label:"Password",tooltip:"NATS password",type:"string",placeholder:"Enter NATS password"},{name:"token",label:"Token",tooltip:"NATS token",type:"string",placeholder:"Enter NATS token"},{name:"tls",label:"TLS",tooltip:"Set to 'on' to enable TLS",type:"on|off"},{name:"tls_skip_verify",label:"TLS Skip Verify",tooltip:'Trust server TLS without verification, defaults to "on" (verify)',type:"on|off"},{name:"ping_interval",label:"Ping Interval",tooltip:"Client ping commands interval in s,m,h,d. Disabled by default",type:"duration",placeholder:"Enter Ping Interval"},{name:"streaming",label:"Streaming",tooltip:"Set to 'on' to use streaming NATS server",type:"on|off"},{name:"streaming_async",label:"Streaming async",tooltip:"Set to 'on' to enable asynchronous publish",type:"on|off"},{name:"streaming_max_pub_acks_in_flight",label:"Streaming max publish ACKS in flight",tooltip:"Number of messages to publish without waiting for ACKs",type:"number",placeholder:"Enter Streaming in flight value"},{name:"streaming_cluster_id",label:"Streaming Cluster ID",tooltip:"Unique ID for NATS streaming cluster",type:"string",placeholder:"Enter Streaming Cluster ID"},{name:"cert_authority",label:"Cert Authority",tooltip:"Path to certificate chain of the target NATS server",type:"string",placeholder:"Enter Cert Authority"},{name:"client_cert",label:"Client Cert",tooltip:"Client cert for NATS mTLS auth",type:"string",placeholder:"Enter Client Cert"},{name:"client_key",label:"Client Key",tooltip:"Client cert key for NATS mTLS authorization",type:"string",placeholder:"Enter Client Key"},...y],[p]:[{name:"url",required:!0,label:"URL",tooltip:"Elasticsearch server's address, with optional authentication info",type:"url",placeholder:"Enter URL"},{name:"index",required:!0,label:"Index",tooltip:"Elasticsearch index to store/update events, index is auto-created",type:"string",placeholder:"Enter Index"},{name:"format",required:!0,label:"Format",tooltip:"'namespace' reflects current bucket/object list and 'access' reflects a journal of object operations, defaults to 'namespace'",type:"enum",placeholder:"Enter Format"},...y],[d]:[{name:"endpoint",required:!0,label:"Endpoint",tooltip:"Webhook server endpoint e.g. http://localhost:8080/minio/events",type:"url",placeholder:"Enter Endpoint"},{name:"auth_token",label:"Auth Token",tooltip:"Opaque string or JWT authorization token",type:"string",placeholder:"Enter auth_token"},...y],[u]:[{name:"nsqd_address",required:!0,label:"NSQD Address",tooltip:"NSQ server address e.g. '127.0.0.1:4150'",type:"address",placeholder:"Enter nsqd_address"},{name:"topic",required:!0,label:"Topic",tooltip:"NSQ topic",type:"string",placeholder:"Enter Topic"},{name:"tls",label:"TLS",tooltip:"Set to 'on' to enable TLS",type:"on|off"},{name:"tls_skip_verify",label:"TLS Skip Verify",tooltip:'Trust server TLS without verification, defaults to "on" (verify)',type:"on|off"},...y]},v={webhook:"notify_webhook",amqp:"notify_amqp",kafka:"notify_kafka",mqtt:"notify_mqtt",nats:"notify_nats",nsq:"notify_nsq",mysql:"notify_mysql",postgresql:"notify_postgres",elasticsearch:"notify_elasticsearch",redis:"notify_redis"},T=e=>v[e]}}]);
//# sourceMappingURL=8017.d5b163f3.chunk.js.map