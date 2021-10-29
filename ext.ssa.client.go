package kyoto

var ssaclient = "<script>(()=>{var z=11;function ne(e,n){var t=n.attributes,r,a,d,s,b;if(!(n.nodeType===z||e.nodeType===z)){for(var S=t.length-1;S>=0;S--)r=t[S],a=r.name,d=r.namespaceURI,s=r.value,d?(a=r.localName||a,b=e.getAttributeNS(d,a),b!==s&&(r.prefix===\"xmlns\"&&(a=r.name),e.setAttributeNS(d,a,s))):(b=e.getAttribute(a),b!==s&&e.setAttribute(a,s));for(var y=e.attributes,w=y.length-1;w>=0;w--)r=y[w],a=r.name,d=r.namespaceURI,d?(a=r.localName||a,n.hasAttributeNS(d,a)||e.removeAttributeNS(d,a)):n.hasAttribute(a)||e.removeAttribute(a)}}var D,ae=\"http://www.w3.org/1999/xhtml\",o=typeof document==\"undefined\"?void 0:document,ie=!!o&&\"content\"in o.createElement(\"template\"),le=!!o&&o.createRange&&\"createContextualFragment\"in o.createRange();function fe(e){var n=o.createElement(\"template\");return n.innerHTML=e,n.content.childNodes[0]}function de(e){D||(D=o.createRange(),D.selectNode(o.body));var n=D.createContextualFragment(e);return n.childNodes[0]}function ue(e){var n=o.createElement(\"body\");return n.innerHTML=e,n.childNodes[0]}function se(e){return e=e.trim(),ie?fe(e):le?de(e):ue(e)}function P(e,n){var t=e.nodeName,r=n.nodeName,a,d;return t===r?!0:(a=t.charCodeAt(0),d=r.charCodeAt(0),a<=90&&d>=97?t===r.toUpperCase():d<=90&&a>=97?r===t.toUpperCase():!1)}function ce(e,n){return!n||n===ae?o.createElement(e):o.createElementNS(n,e)}function oe(e,n){for(var t=e.firstChild;t;){var r=t.nextSibling;n.appendChild(t),t=r}return n}function I(e,n,t){e[t]!==n[t]&&(e[t]=n[t],e[t]?e.setAttribute(t,\"\"):e.removeAttribute(t))}var j={OPTION:function(e,n){var t=e.parentNode;if(t){var r=t.nodeName.toUpperCase();r===\"OPTGROUP\"&&(t=t.parentNode,r=t&&t.nodeName.toUpperCase()),r===\"SELECT\"&&!t.hasAttribute(\"multiple\")&&(e.hasAttribute(\"selected\")&&!n.selected&&(e.setAttribute(\"selected\",\"selected\"),e.removeAttribute(\"selected\")),t.selectedIndex=-1)}I(e,n,\"selected\")},INPUT:function(e,n){I(e,n,\"checked\"),I(e,n,\"disabled\"),e.value!==n.value&&(e.value=n.value),n.hasAttribute(\"value\")||e.removeAttribute(\"value\")},TEXTAREA:function(e,n){var t=n.value;e.value!==t&&(e.value=t);var r=e.firstChild;if(r){var a=r.nodeValue;if(a==t||!t&&a==e.placeholder)return;r.nodeValue=t}},SELECT:function(e,n){if(!n.hasAttribute(\"multiple\")){for(var t=-1,r=0,a=e.firstChild,d,s;a;)if(s=a.nodeName&&a.nodeName.toUpperCase(),s===\"OPTGROUP\")d=a,a=d.firstChild;else{if(s===\"OPTION\"){if(a.hasAttribute(\"selected\")){t=r;break}r++}a=a.nextSibling,!a&&d&&(a=d.nextSibling,d=null)}e.selectedIndex=t}}},N=1,ve=11,k=3,W=8;function A(){}function he(e){if(e)return e.getAttribute&&e.getAttribute(\"id\")||e.id}function ge(e){return function(t,r,a){if(a||(a={}),typeof r==\"string\")if(t.nodeName===\"#document\"||t.nodeName===\"HTML\"||t.nodeName===\"BODY\"){var d=r;r=o.createElement(\"html\"),r.innerHTML=d}else r=se(r);var s=a.getNodeKey||he,b=a.onBeforeNodeAdded||A,S=a.onNodeAdded||A,y=a.onBeforeElUpdated||A,w=a.onElUpdated||A,Q=a.onBeforeNodeDiscarded||A,O=a.onNodeDiscarded||A,Z=a.onBeforeElChildrenUpdated||A,B=a.childrenOnly===!0,T=Object.create(null),L=[];function M(f){L.push(f)}function X(f,l){if(f.nodeType===N)for(var i=f.firstChild;i;){var u=void 0;l&&(u=s(i))?M(u):(O(i),i.firstChild&&X(i,l)),i=i.nextSibling}}function x(f,l,i){Q(f)!==!1&&(l&&l.removeChild(f),O(f),X(f,i))}function G(f){if(f.nodeType===N||f.nodeType===ve)for(var l=f.firstChild;l;){var i=s(l);i&&(T[i]=l),G(l),l=l.nextSibling}}G(t);function C(f){S(f);for(var l=f.firstChild;l;){var i=l.nextSibling,u=s(l);if(u){var v=T[u];v&&P(l,v)?(l.parentNode.replaceChild(v,l),E(v,l)):C(l)}else C(l);l=i}}function ee(f,l,i){for(;l;){var u=l.nextSibling;(i=s(l))?M(i):x(l,f,!0),l=u}}function E(f,l,i){var u=s(l);u&&delete T[u],!(!i&&(y(f,l)===!1||(e(f,l),w(f),Z(f,l)===!1)))&&(f.nodeName!==\"TEXTAREA\"?te(f,l):j.TEXTAREA(f,l))}function te(f,l){var i=l.firstChild,u=f.firstChild,v,h,m,U,g;e:for(;i;){for(U=i.nextSibling,v=s(i);u;){if(m=u.nextSibling,i.isSameNode&&i.isSameNode(u)){i=U,u=m;continue e}h=s(u);var R=u.nodeType,p=void 0;if(R===i.nodeType&&(R===N?(v?v!==h&&((g=T[v])?m===g?p=!1:(f.insertBefore(g,u),h?M(h):x(u,f,!0),u=g):p=!1):h&&(p=!1),p=p!==!1&&P(u,i),p&&E(u,i)):(R===k||R==W)&&(p=!0,u.nodeValue!==i.nodeValue&&(u.nodeValue=i.nodeValue))),p){i=U,u=m;continue e}h?M(h):x(u,f,!0),u=m}if(v&&(g=T[v])&&P(g,i))f.appendChild(g),E(g,i);else{var $=b(i);$!==!1&&($&&(i=$),i.actualize&&(i=i.actualize(f.ownerDocument||o)),f.appendChild(i),C(i))}i=U,u=m}ee(f,u,h);var K=j[f.nodeName];K&&K(f,l)}var c=t,H=c.nodeType,J=r.nodeType;if(!B){if(H===N)J===N?P(t,r)||(O(t),c=oe(t,ce(r.nodeName,r.namespaceURI))):c=r;else if(H===k||H===W){if(J===H)return c.nodeValue!==r.nodeValue&&(c.nodeValue=r.nodeValue),c;c=r}}if(c===r)O(t);else{if(r.isSameNode&&r.isSameNode(c))return;if(E(c,r,B),L)for(var V=0,re=L.length;V<re;V++){var F=T[L[V]];F&&x(F,F.parentNode,!1)}}return!B&&c!==t&&t.parentNode&&(c.actualize&&(c=c.actualize(t.ownerDocument||o)),t.parentNode.replaceChild(c,t)),c}}var pe=ge(ne),Y=pe;function _(e){let n=e.starter;if(e.id){let t=document.getElementById(e.id);if(!t)throw new Error(`Error while locating root: can't find direct with ${e}`);n=t}else{let t=0;for(;;){if(!n.parentElement)throw new Error(`Error while locating root: can't find parent with ${e}`);if(!n.getAttribute(\"state\"))n=n.parentElement;else if(e.depth&&t!=e.depth)n=n.parentElement,t++;else break}}return n}function Ae(e){return e.includes(\":\")&&(e=e.split(\":\")[1]),e.includes(\"$\")&&(e=e.replaceAll(\"$\",\"\")),e}function q(e,n,...t){let r=_({starter:e,depth:n.split(\"\").filter(d=>d===\"$\").length,id:n.includes(\":\")?n.split(\":\")[0]:void 0}),a=new FormData;a.set(\"State\",r.getAttribute(\"state\")||\"{}\"),a.set(\"Args\",JSON.stringify(t)),fetch(`/SSA/${r.getAttribute(\"name\")}/${Ae(n)}`,{method:\"POST\",body:a}).then(d=>d.headers.get(\"X-Redirect\")?(window.location.href=d.headers.get(\"X-Redirect\"),\"\"):d.text()).then(d=>{if(!!d){if(r.hasAttribute(\"ssa-replace\")){r.outerHTML=d;return}try{Y(r,d)}catch(s){console.log(\"Fallback from morphdom to root.outerHTML due to error\",s),r.outerHTML=d}}}).catch(d=>{console.log(d)})}function be(e,n){let t=_({starter:e,depth:n.split(\"\").filter(a=>a===\"$\").length,id:n.includes(\":\")?n.split(\":\")[0]:void 0});if(!t.getAttribute(\"state\"))throw new Error(\"Bind call error: component state is underfined\");let r=JSON.parse(atob(t.getAttribute(\"state\")));r[n]=e.value,t.setAttribute(\"state\",btoa(JSON.stringify(r)))}function Te(e,n){n.preventDefault();let t=_({starter:e});if(!t.getAttribute(\"state\"))throw new Error(\"Bind call error: component state is underfined\");let r=JSON.parse(atob(t.getAttribute(\"state\"))),a=new FormData(n.target),d=Object.fromEntries(a.entries());return Object.entries(d).forEach(s=>{r[s[0]]=s[1]}),t.setAttribute(\"state\",btoa(JSON.stringify(r))),q(t,\"Submit\"),!1}window._LocaleRoot=_;window.Action=q;window.Bind=be;window.FormSubmit=Te;})();</script>"
