package kyoto

var ssaclient = "<script>(()=>{var z=11;function ne(e,n){var t=n.attributes,r,a,s,u,b;if(!(n.nodeType===z||e.nodeType===z)){for(var S=t.length-1;S>=0;S--)r=t[S],a=r.name,s=r.namespaceURI,u=r.value,s?(a=r.localName||a,b=e.getAttributeNS(s,a),b!==u&&(r.prefix===\"xmlns\"&&(a=r.name),e.setAttributeNS(s,a,u))):(b=e.getAttribute(a),b!==u&&e.setAttribute(a,u));for(var O=e.attributes,w=O.length-1;w>=0;w--)r=O[w],a=r.name,s=r.namespaceURI,s?(a=r.localName||a,n.hasAttributeNS(s,a)||e.removeAttributeNS(s,a)):n.hasAttribute(a)||e.removeAttribute(a)}}var _,ae=\"http://www.w3.org/1999/xhtml\",c=typeof document==\"undefined\"?void 0:document,ie=!!c&&\"content\"in c.createElement(\"template\"),le=!!c&&c.createRange&&\"createContextualFragment\"in c.createRange();function de(e){var n=c.createElement(\"template\");return n.innerHTML=e,n.content.childNodes[0]}function se(e){_||(_=c.createRange(),_.selectNode(c.body));var n=_.createContextualFragment(e);return n.childNodes[0]}function fe(e){var n=c.createElement(\"body\");return n.innerHTML=e,n.childNodes[0]}function ue(e){return e=e.trim(),ie?de(e):le?se(e):fe(e)}function P(e,n){var t=e.nodeName,r=n.nodeName,a,s;return t===r?!0:(a=t.charCodeAt(0),s=r.charCodeAt(0),a<=90&&s>=97?t===r.toUpperCase():s<=90&&a>=97?r===t.toUpperCase():!1)}function oe(e,n){return!n||n===ae?c.createElement(e):c.createElementNS(n,e)}function ce(e,n){for(var t=e.firstChild;t;){var r=t.nextSibling;n.appendChild(t),t=r}return n}function I(e,n,t){e[t]!==n[t]&&(e[t]=n[t],e[t]?e.setAttribute(t,\"\"):e.removeAttribute(t))}var j={OPTION:function(e,n){var t=e.parentNode;if(t){var r=t.nodeName.toUpperCase();r===\"OPTGROUP\"&&(t=t.parentNode,r=t&&t.nodeName.toUpperCase()),r===\"SELECT\"&&!t.hasAttribute(\"multiple\")&&(e.hasAttribute(\"selected\")&&!n.selected&&(e.setAttribute(\"selected\",\"selected\"),e.removeAttribute(\"selected\")),t.selectedIndex=-1)}I(e,n,\"selected\")},INPUT:function(e,n){I(e,n,\"checked\"),I(e,n,\"disabled\"),e.value!==n.value&&(e.value=n.value),n.hasAttribute(\"value\")||e.removeAttribute(\"value\")},TEXTAREA:function(e,n){var t=n.value;e.value!==t&&(e.value=t);var r=e.firstChild;if(r){var a=r.nodeValue;if(a==t||!t&&a==e.placeholder)return;r.nodeValue=t}},SELECT:function(e,n){if(!n.hasAttribute(\"multiple\")){for(var t=-1,r=0,a=e.firstChild,s,u;a;)if(u=a.nodeName&&a.nodeName.toUpperCase(),u===\"OPTGROUP\")s=a,a=s.firstChild;else{if(u===\"OPTION\"){if(a.hasAttribute(\"selected\")){t=r;break}r++}a=a.nextSibling,!a&&s&&(a=s.nextSibling,s=null)}e.selectedIndex=t}}},y=1,ve=11,k=3,W=8;function A(){}function he(e){if(e)return e.getAttribute&&e.getAttribute(\"id\")||e.id}function pe(e){return function(t,r,a){if(a||(a={}),typeof r==\"string\")if(t.nodeName===\"#document\"||t.nodeName===\"HTML\"||t.nodeName===\"BODY\"){var s=r;r=c.createElement(\"html\"),r.innerHTML=s}else r=ue(r);var u=a.getNodeKey||he,b=a.onBeforeNodeAdded||A,S=a.onNodeAdded||A,O=a.onBeforeElUpdated||A,w=a.onElUpdated||A,Q=a.onBeforeNodeDiscarded||A,L=a.onNodeDiscarded||A,Z=a.onBeforeElChildrenUpdated||A,B=a.childrenOnly===!0,T=Object.create(null),M=[];function x(d){M.push(d)}function X(d,l){if(d.nodeType===y)for(var i=d.firstChild;i;){var f=void 0;l&&(f=u(i))?x(f):(L(i),i.firstChild&&X(i,l)),i=i.nextSibling}}function E(d,l,i){Q(d)!==!1&&(l&&l.removeChild(d),L(d),X(d,i))}function G(d){if(d.nodeType===y||d.nodeType===ve)for(var l=d.firstChild;l;){var i=u(l);i&&(T[i]=l),G(l),l=l.nextSibling}}G(t);function C(d){S(d);for(var l=d.firstChild;l;){var i=l.nextSibling,f=u(l);if(f){var v=T[f];v&&P(l,v)?(l.parentNode.replaceChild(v,l),H(v,l)):C(l)}else C(l);l=i}}function ee(d,l,i){for(;l;){var f=l.nextSibling;(i=u(l))?x(i):E(l,d,!0),l=f}}function H(d,l,i){var f=u(l);f&&delete T[f],!(!i&&(O(d,l)===!1||(e(d,l),w(d),Z(d,l)===!1)))&&(d.nodeName!==\"TEXTAREA\"?te(d,l):j.TEXTAREA(d,l))}function te(d,l){var i=l.firstChild,f=d.firstChild,v,h,m,R,p;e:for(;i;){for(R=i.nextSibling,v=u(i);f;){if(m=f.nextSibling,i.isSameNode&&i.isSameNode(f)){i=R,f=m;continue e}h=u(f);var D=f.nodeType,g=void 0;if(D===i.nodeType&&(D===y?(v?v!==h&&((p=T[v])?m===p?g=!1:(d.insertBefore(p,f),h?x(h):E(f,d,!0),f=p):g=!1):h&&(g=!1),g=g!==!1&&P(f,i),g&&H(f,i)):(D===k||D==W)&&(g=!0,f.nodeValue!==i.nodeValue&&(f.nodeValue=i.nodeValue))),g){i=R,f=m;continue e}h?x(h):E(f,d,!0),f=m}if(v&&(p=T[v])&&P(p,i))d.appendChild(p),H(p,i);else{var $=b(i);$!==!1&&($&&(i=$),i.actualize&&(i=i.actualize(d.ownerDocument||c)),d.appendChild(i),C(i))}i=R,f=m}ee(d,f,h);var K=j[d.nodeName];K&&K(d,l)}var o=t,U=o.nodeType,J=r.nodeType;if(!B){if(U===y)J===y?P(t,r)||(L(t),o=ce(t,oe(r.nodeName,r.namespaceURI))):o=r;else if(U===k||U===W){if(J===U)return o.nodeValue!==r.nodeValue&&(o.nodeValue=r.nodeValue),o;o=r}}if(o===r)L(t);else{if(r.isSameNode&&r.isSameNode(o))return;if(H(o,r,B),M)for(var V=0,re=M.length;V<re;V++){var F=T[M[V]];F&&E(F,F.parentNode,!1)}}return!B&&o!==t&&t.parentNode&&(o.actualize&&(o=o.actualize(t.ownerDocument||c)),t.parentNode.replaceChild(o,t)),o}}var ge=pe(ne),q=ge;function N(e){let n=e.starter;if(e.id){let t=document.getElementById(e.id);if(!t)throw new Error(`Error while locating root: can't find direct with ${e}`);n=t}else{let t=0;for(;;){if(!n.parentElement)throw new Error(`Error while locating root: can't find parent with ${e}`);if(!n.getAttribute(\"state\"))n=n.parentElement;else if(e.depth&&t!=e.depth)n=n.parentElement,t++;else break}}return n}function Ae(e){return e.includes(\":\")&&(e=e.split(\":\")[1]),e.includes(\"$\")&&(e=e.replaceAll(\"$\",\"\")),e}function Y(e,n,...t){let r=N({starter:e,depth:n.split(\"\").filter(s=>s===\"$\").length,id:n.includes(\":\")?n.split(\":\")[0]:void 0});me(e);let a=new FormData;a.set(\"State\",r.getAttribute(\"state\")||\"{}\"),a.set(\"Args\",JSON.stringify(t)),fetch(`/SSA/${r.getAttribute(\"name\")}/${Ae(n)}`,{method:\"POST\",body:a}).then(s=>s.headers.get(\"X-Redirect\")?(window.location.href=s.headers.get(\"X-Redirect\"),\"\"):s.text()).then(s=>{if(!!s){if(r.hasAttribute(\"ssa:replace\")){r.outerHTML=s;return}try{q(r,s)}catch(u){console.log(\"Fallback from morphdom to root.outerHTML due to error\",u),r.outerHTML=s}}}).catch(s=>{console.log(s)})}function be(e,n){let t=N({starter:e,depth:n.split(\"\").filter(a=>a===\"$\").length,id:n.includes(\":\")?n.split(\":\")[0]:void 0});if(!t.getAttribute(\"state\"))throw new Error(\"Bind call error: component state is underfined\");let r=JSON.parse(atob(t.getAttribute(\"state\")));r[n]=e.value,t.setAttribute(\"state\",btoa(JSON.stringify(r)))}function Te(e,n){n.preventDefault();let t=N({starter:e});if(!t.getAttribute(\"state\"))throw new Error(\"Bind call error: component state is underfined\");let r=JSON.parse(atob(t.getAttribute(\"state\"))),a=new FormData(n.target),s=Object.fromEntries(a.entries());return Object.entries(s).forEach(u=>{r[u[0]]=u[1]}),t.setAttribute(\"state\",btoa(JSON.stringify(r))),Y(t,\"Submit\"),!1}function me(e){N({starter:e}).querySelectorAll(\"[ssa\\\\:oncall\\\\.display]\").forEach(r=>{let a=r.getAttribute(\"ssa:oncall.display\");a!=\"\"&&r.setAttribute(\"style\",\"display: \"+a)})}window._LocaleRoot=N;window.Action=Y;window.Bind=be;window.FormSubmit=Te;})();</script>"
