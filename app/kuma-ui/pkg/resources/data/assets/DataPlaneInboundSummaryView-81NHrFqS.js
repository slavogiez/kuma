import{d as R,l as x,V as D,a as l,o as t,b as m,w as s,f as c,F as g,t as v,c as h,m as y,q as N,e as _,C as V,x as S,y as B,a2 as T,_ as A}from"./index-Co5ZlPWJ.js";import{N as F}from"./NavTabs-TBb2WjpU.js";const P=n=>(S("data-v-26827d66"),n=n(),B(),n),$={class:"summary-title-wrapper"},q=P(()=>y("img",{"aria-hidden":"true",src:T},null,-1)),E={class:"summary-title"},j=R({__name:"DataPlaneInboundSummaryView",props:{dataplaneType:{},gateway:{},inbounds:{}},setup(n){var w;const{t:C}=x(),I=D(),o=n,k=(((w=I.getRoutes().find(e=>e.name==="data-plane-inbound-summary-view"))==null?void 0:w.children)??[]).map(e=>{var i,a;const p=typeof e.name>"u"?(i=e.children)==null?void 0:i[0]:e,r=p.name,u=((a=p.meta)==null?void 0:a.module)??"";return{title:C(`data-planes.routes.item.navigation.${r}`),routeName:r,module:u}});return(e,p)=>{const r=l("DataCollection"),u=l("RouterView"),f=l("AppView"),i=l("RouteView");return t(),m(i,{name:"data-plane-inbound-summary-view",params:{service:""}},{default:s(({route:a})=>[_(f,null,{title:s(()=>[y("div",$,[q,c(),y("h2",E,[o.gateway?(t(),h(g,{key:0},[c(v(a.params.service),1)],64)):(t(),h(g,{key:1},[c(`
              Inbound :`+v(a.params.service.replace("localhost_","")),1)],64))])])]),default:s(()=>[c(),_(F,{tabs:N(k)},null,8,["tabs"]),c(),_(u,null,{default:s(b=>[o.gateway?(t(),m(V(b.Component),{key:0,gateway:o.gateway},null,8,["gateway"])):(t(),m(r,{key:1,items:o.inbounds,predicate:d=>`localhost_${d.port}`===a.params.service,find:!0},{default:s(({items:d})=>[(t(),m(V(b.Component),{inbound:d[0],gateway:o.gateway},null,8,["inbound","gateway"]))]),_:2},1032,["items","predicate"]))]),_:2},1024)]),_:2},1024)]),_:1})}}}),J=A(j,[["__scopeId","data-v-26827d66"]]);export{J as default};
