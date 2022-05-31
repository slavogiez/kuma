(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["mesh-gateway-routes"],{9670:function(e,t,a){"use strict";a.r(t);var n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"meshgatewayroutes relative"},[a("DocumentationLink",{attrs:{href:e.docsURL}}),a("div",{staticClass:"mb-4"},[a("KAlert",{attrs:{appearance:"warning"},scopedSlots:e._u([{key:"alertMessage",fn:function(){return[a("p",[a("strong",[e._v("Warning")]),e._v(" This policy is experimental. If you encountered any problem please open an "),a("a",{attrs:{href:"https://github.com/kumahq/kuma/issues/new/choose",target:"_blank",rel:"noopener noreferrer"}},[e._v("issue")])])]},proxy:!0}])})],1),a("FrameSkeleton",[a("DataOverview",{attrs:{"page-size":e.pageSize,"has-error":e.hasError,"is-loading":e.isLoading,"empty-state":e.empty_state,"table-data":e.tableData,"table-data-is-empty":e.isEmpty,next:e.next},on:{tableAction:e.tableAction,loadData:function(t){return e.loadData(t)}},scopedSlots:e._u([{key:"additionalControls",fn:function(){return[e.$route.query.ns?a("KButton",{staticClass:"back-button",attrs:{appearance:"primary",size:"small",to:{name:"meshgatewayroutes"}}},[a("span",{staticClass:"custom-control-icon"},[e._v(" ← ")]),e._v(" View All ")]):e._e()]},proxy:!0}])},[e._v(" > ")]),!1===e.isEmpty?a("Tabs",{attrs:{"has-error":e.hasError,"is-loading":e.isLoading,tabs:e.tabs,"initial-tab-override":"overview"},scopedSlots:e._u([{key:"tabHeader",fn:function(){return[a("div",[a("h3",[e._v(" Gateway Route: "+e._s(e.entity.name)+" ")])]),a("div",[a("EntityURLControl",{attrs:{name:e.entity.name,mesh:e.entity.mesh}})],1)]},proxy:!0},{key:"overview",fn:function(){return[a("LabelList",[a("div",[a("ul",e._l(e.entity,(function(t,n){return a("li",{key:n},[a("h4",[e._v(e._s(n))]),a("p",[e._v(" "+e._s(t)+" ")])])})),0)])])]},proxy:!0},{key:"affected-dpps",fn:function(){return[a("PolicyConnections",{attrs:{mesh:e.rawEntity.mesh,"policy-name":e.rawEntity.name,"policy-type":"meshgatewayroutes"}})]},proxy:!0},{key:"yaml",fn:function(){return[a("YamlView",{attrs:{lang:"yaml",content:e.rawEntity}})]},proxy:!0}],null,!1,2048740073)}):e._e()],1)],1)},s=[],i=(a("7db0"),a("b0c0"),a("96cf"),a("c964")),r=a("f3f3"),o=a("2f62"),l=a("bc1e"),c=a("1d3a"),u=a("0f82"),y=a("6663"),m=a("1d10"),h=a("2778"),p=a("251b"),d=a("14eb"),f=a("ff9d"),b=a("0ada"),g=a("6524"),v=a("c6ec"),w={name:"MeshGatewayRoutes",metaInfo:{title:"Mesh Gateway Routes"},components:{EntityURLControl:y["a"],FrameSkeleton:m["a"],DataOverview:h["a"],Tabs:p["a"],YamlView:f["a"],LabelList:b["a"],PolicyConnections:d["a"],DocumentationLink:g["a"]},data:function(){return{isLoading:!0,isEmpty:!1,hasError:!1,empty_state:{title:"No Data",message:"There are no Mesh Gateway Routes present."},tableData:{headers:[{key:"actions",hideLabel:!0},{label:"Name",key:"name"},{label:"Mesh",key:"mesh"},{label:"Type",key:"type"}],data:[]},tabs:[{hash:"#overview",title:"Overview"},{hash:"#affected-dpps",title:"Affected DPPs"},{hash:"#yaml",title:"YAML"}],entity:{},rawData:[],rawEntity:{},pageSize:v["g"],next:null}},computed:Object(r["a"])(Object(r["a"])({},Object(o["c"])({version:"config/getVersion"})),{},{docsURL:function(){return"https://kuma.io/docs/".concat(this.version,"/policies/meshgatewayroute/")}}),watch:{$route:function(e,t){this.init()}},beforeMount:function(){this.init()},methods:{init:function(){this.loadData()},tableAction:function(e){var t=e;this.getEntity(t)},loadData:function(){var e=arguments,t=this;return Object(i["a"])(regeneratorRuntime.mark((function a(){var n,s,i,r,o,l;return regeneratorRuntime.wrap((function(a){while(1)switch(a.prev=a.next){case 0:return n=e.length>0&&void 0!==e[0]?e[0]:"0",t.isLoading=!0,s=t.$route.query.ns||null,i=t.$route.params.mesh||null,a.prev=4,a.next=7,Object(c["a"])({getSingleEntity:u["a"].getMeshGatewayRoute.bind(u["a"]),getAllEntities:u["a"].getAllMeshGatewayRoutes.bind(u["a"]),getAllEntitiesFromMesh:u["a"].getAllMeshGatewayRoutesFromMesh.bind(u["a"]),mesh:i,query:s,size:t.pageSize,offset:n});case 7:r=a.sent,o=r.data,l=r.next,t.next=l,o.length?(t.isEmpty=!1,t.rawData=o,t.tableData.data=o,t.getEntity({name:o[0].name})):(t.tableData.data=[],t.isEmpty=!0),a.next=19;break;case 14:a.prev=14,a.t0=a["catch"](4),t.hasError=!0,t.isEmpty=!0,console.error(a.t0);case 19:return a.prev=19,t.isLoading=!1,a.finish(19);case 22:case"end":return a.stop()}}),a,null,[[4,14,19,22]])})))()},getEntity:function(e){var t=["type","name","mesh"],a=this.rawData.find((function(t){return t.name===e.name}));this.rawEntity=Object(l["j"])(a),this.entity=Object(l["d"])(a,t)}}},k=w,E=a("2877"),_=Object(E["a"])(k,n,s,!1,null,null,null);t["default"]=_.exports}}]);