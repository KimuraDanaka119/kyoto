package kyoto

var ssaclient = "<script>(()=>{var j=11;function ne(e,r){var t=r.attributes,n,a,s,f,b;if(!(r.nodeType===j||e.nodeType===j)){for(var S=t.length-1;S>=0;S--)n=t[S],a=n.name,s=n.namespaceURI,f=n.value,s?(a=n.localName||a,b=e.getAttributeNS(s,a),b!==f&&(n.prefix===\"xmlns\"&&(a=n.name),e.setAttributeNS(s,a,f))):(b=e.getAttribute(a),b!==f&&e.setAttribute(a,f));for(var L=e.attributes,w=L.length-1;w>=0;w--)n=L[w],a=n.name,s=n.namespaceURI,s?(a=n.localName||a,r.hasAttributeNS(s,a)||e.removeAttributeNS(s,a)):r.hasAttribute(a)||e.removeAttribute(a)}}var _,ae=\"http://www.w3.org/1999/xhtml\",c=typeof document==\"undefined\"?void 0:document,ie=!!c&&\"content\"in c.createElement(\"template\"),le=!!c&&c.createRange&&\"createContextualFragment\"in c.createRange();function de(e){var r=c.createElement(\"template\");return r.innerHTML=e,r.content.childNodes[0]}function se(e){_||(_=c.createRange(),_.selectNode(c.body));var r=_.createContextualFragment(e);return r.childNodes[0]}function ue(e){var r=c.createElement(\"body\");return r.innerHTML=e,r.childNodes[0]}function fe(e){return e=e.trim(),ie?de(e):le?se(e):ue(e)}function P(e,r){var t=e.nodeName,n=r.nodeName,a,s;return t===n?!0:(a=t.charCodeAt(0),s=n.charCodeAt(0),a<=90&&s>=97?t===n.toUpperCase():s<=90&&a>=97?n===t.toUpperCase():!1)}function oe(e,r){return!r||r===ae?c.createElement(e):c.createElementNS(r,e)}function ce(e,r){for(var t=e.firstChild;t;){var n=t.nextSibling;r.appendChild(t),t=n}return r}function I(e,r,t){e[t]!==r[t]&&(e[t]=r[t],e[t]?e.setAttribute(t,\"\"):e.removeAttribute(t))}var k={OPTION:function(e,r){var t=e.parentNode;if(t){var n=t.nodeName.toUpperCase();n===\"OPTGROUP\"&&(t=t.parentNode,n=t&&t.nodeName.toUpperCase()),n===\"SELECT\"&&!t.hasAttribute(\"multiple\")&&(e.hasAttribute(\"selected\")&&!r.selected&&(e.setAttribute(\"selected\",\"selected\"),e.removeAttribute(\"selected\")),t.selectedIndex=-1)}I(e,r,\"selected\")},INPUT:function(e,r){I(e,r,\"checked\"),I(e,r,\"disabled\"),e.value!==r.value&&(e.value=r.value),r.hasAttribute(\"value\")||e.removeAttribute(\"value\")},TEXTAREA:function(e,r){var t=r.value;e.value!==t&&(e.value=t);var n=e.firstChild;if(n){var a=n.nodeValue;if(a==t||!t&&a==e.placeholder)return;n.nodeValue=t}},SELECT:function(e,r){if(!r.hasAttribute(\"multiple\")){for(var t=-1,n=0,a=e.firstChild,s,f;a;)if(f=a.nodeName&&a.nodeName.toUpperCase(),f===\"OPTGROUP\")s=a,a=s.firstChild;else{if(f===\"OPTION\"){if(a.hasAttribute(\"selected\")){t=n;break}n++}a=a.nextSibling,!a&&s&&(a=s.nextSibling,s=null)}e.selectedIndex=t}}},y=1,ve=11,q=3,W=8;function A(){}function he(e){if(e)return e.getAttribute&&e.getAttribute(\"id\")||e.id}function pe(e){return function(t,n,a){if(a||(a={}),typeof n==\"string\")if(t.nodeName===\"#document\"||t.nodeName===\"HTML\"||t.nodeName===\"BODY\"){var s=n;n=c.createElement(\"html\"),n.innerHTML=s}else n=fe(n);var f=a.getNodeKey||he,b=a.onBeforeNodeAdded||A,S=a.onNodeAdded||A,L=a.onBeforeElUpdated||A,w=a.onElUpdated||A,Q=a.onBeforeNodeDiscarded||A,O=a.onNodeDiscarded||A,Z=a.onBeforeElChildrenUpdated||A,B=a.childrenOnly===!0,T=Object.create(null),M=[];function E(d){M.push(d)}function G(d,l){if(d.nodeType===y)for(var i=d.firstChild;i;){var u=void 0;l&&(u=f(i))?E(u):(O(i),i.firstChild&&G(i,l)),i=i.nextSibling}}function x(d,l,i){Q(d)!==!1&&(l&&l.removeChild(d),O(d),G(d,i))}function J(d){if(d.nodeType===y||d.nodeType===ve)for(var l=d.firstChild;l;){var i=f(l);i&&(T[i]=l),J(l),l=l.nextSibling}}J(t);function C(d){S(d);for(var l=d.firstChild;l;){var i=l.nextSibling,u=f(l);if(u){var v=T[u];v&&P(l,v)?(l.parentNode.replaceChild(v,l),H(v,l)):C(l)}else C(l);l=i}}function ee(d,l,i){for(;l;){var u=l.nextSibling;(i=f(l))?E(i):x(l,d,!0),l=u}}function H(d,l,i){var u=f(l);u&&delete T[u],!(!i&&(L(d,l)===!1||(e(d,l),w(d),Z(d,l)===!1)))&&(d.nodeName!==\"TEXTAREA\"?te(d,l):k.TEXTAREA(d,l))}function te(d,l){var i=l.firstChild,u=d.firstChild,v,h,m,D,p;e:for(;i;){for(D=i.nextSibling,v=f(i);u;){if(m=u.nextSibling,i.isSameNode&&i.isSameNode(u)){i=D,u=m;continue e}h=f(u);var R=u.nodeType,g=void 0;if(R===i.nodeType&&(R===y?(v?v!==h&&((p=T[v])?m===p?g=!1:(d.insertBefore(p,u),h?E(h):x(u,d,!0),u=p):g=!1):h&&(g=!1),g=g!==!1&&P(u,i),g&&H(u,i)):(R===q||R==W)&&(g=!0,u.nodeValue!==i.nodeValue&&(u.nodeValue=i.nodeValue))),g){i=D,u=m;continue e}h?E(h):x(u,d,!0),u=m}if(v&&(p=T[v])&&P(p,i))d.appendChild(p),H(p,i);else{var $=b(i);$!==!1&&($&&(i=$),i.actualize&&(i=i.actualize(d.ownerDocument||c)),d.appendChild(i),C(i))}i=D,u=m}ee(d,u,h);var z=k[d.nodeName];z&&z(d,l)}var o=t,U=o.nodeType,K=n.nodeType;if(!B){if(U===y)K===y?P(t,n)||(O(t),o=ce(t,oe(n.nodeName,n.namespaceURI))):o=n;else if(U===q||U===W){if(K===U)return o.nodeValue!==n.nodeValue&&(o.nodeValue=n.nodeValue),o;o=n}}if(o===n)O(t);else{if(n.isSameNode&&n.isSameNode(o))return;if(H(o,n,B),M)for(var V=0,re=M.length;V<re;V++){var F=T[M[V]];F&&x(F,F.parentNode,!1)}}return!B&&o!==t&&t.parentNode&&(o.actualize&&(o=o.actualize(t.ownerDocument||c)),t.parentNode.replaceChild(o,t)),o}}var ge=pe(ne),Y=ge;function N(e){let r=e.starter;if(e.id){let t=document.getElementById(e.id);if(!t)throw new Error(`Error while locating root: can't find direct with ${e}`);r=t}else{let t=0;for(;;){if(!r.parentElement)throw new Error(`Error while locating root: can't find parent with ${e}`);if(!r.getAttribute(\"state\"))r=r.parentElement;else if(e.depth&&t!=e.depth)r=r.parentElement,t++;else break}}return r}function Ae(e){return e.includes(\":\")&&(e=e.split(\":\")[1]),e.includes(\"$\")&&(e=e.replaceAll(\"$\",\"\")),e}function be(e){N({starter:e}).querySelectorAll(\"[ssa\\\\:oncall\\\\.display]\").forEach(n=>{let a=n.getAttribute(\"ssa:oncall.display\");a!=\"\"&&n.setAttribute(\"style\",\"display: \"+a)})}function Te(){document.querySelectorAll(\"[ssa\\\\:onload]\").forEach(e=>{let r=e.getAttribute(\"ssa:onload\");r&&r!=\"\"&&X(e,r)})}function X(e,r,...t){let n=N({starter:e,depth:r.split(\"\").filter(s=>s===\"$\").length,id:r.includes(\":\")?r.split(\":\")[0]:void 0});be(e);let a=new FormData;a.set(\"State\",n.getAttribute(\"state\")||\"{}\"),a.set(\"Args\",JSON.stringify(t)),fetch(`/SSA/${n.getAttribute(\"name\")}/${Ae(r)}`,{method:\"POST\",body:a}).then(s=>s.headers.get(\"X-Redirect\")?(window.location.href=s.headers.get(\"X-Redirect\"),\"\"):s.text()).then(s=>{if(!!s){if(n.hasAttribute(\"ssa:replace\")){n.outerHTML=s;return}try{Y(n,s)}catch(f){console.log(\"Fallback from morphdom to root.outerHTML due to error\",f),n.outerHTML=s}}}).catch(s=>{console.log(s)})}function me(e,r){let t=N({starter:e,depth:r.split(\"\").filter(a=>a===\"$\").length,id:r.includes(\":\")?r.split(\":\")[0]:void 0});if(!t.getAttribute(\"state\"))throw new Error(\"Bind call error: component state is underfined\");let n=JSON.parse(atob(t.getAttribute(\"state\")));n[r]=e.value,t.setAttribute(\"state\",btoa(JSON.stringify(n)))}function Se(e,r){r.preventDefault();let t=N({starter:e});if(!t.getAttribute(\"state\"))throw new Error(\"Bind call error: component state is underfined\");let n=JSON.parse(atob(t.getAttribute(\"state\"))),a=new FormData(r.target),s=Object.fromEntries(a.entries());return Object.entries(s).forEach(f=>{n[f[0]]=f[1]}),t.setAttribute(\"state\",btoa(JSON.stringify(n))),X(t,\"Submit\"),!1}window._LocaleRoot=N;window.Action=X;window.Bind=me;window.FormSubmit=Se;document.addEventListener(\"DOMContentLoaded\",Te);})();</script>"
