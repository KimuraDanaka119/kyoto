package kyoto

var ActionClient = "<script>(()=>{var z=11;function ae(e,r){var t=r.attributes,n,a,i,d,v;if(!(r.nodeType===z||e.nodeType===z)){for(var f=t.length-1;f>=0;f--)n=t[f],a=n.name,i=n.namespaceURI,d=n.value,i?(a=n.localName||a,v=e.getAttributeNS(i,a),v!==d&&(n.prefix===\"xmlns\"&&(a=n.name),e.setAttributeNS(i,a,d))):(v=e.getAttribute(a),v!==d&&e.setAttribute(a,d));for(var b=e.attributes,p=b.length-1;p>=0;p--)n=b[p],a=n.name,i=n.namespaceURI,i?(a=n.localName||a,r.hasAttributeNS(i,a)||e.removeAttributeNS(i,a)):r.hasAttribute(a)||e.removeAttribute(a)}}var R,ie=\"http://www.w3.org/1999/xhtml\",h=typeof document==\"undefined\"?void 0:document,le=!!h&&\"content\"in h.createElement(\"template\"),oe=!!h&&h.createRange&&\"createContextualFragment\"in h.createRange();function se(e){var r=h.createElement(\"template\");return r.innerHTML=e,r.content.childNodes[0]}function de(e){R||(R=h.createRange(),R.selectNode(h.body));var r=R.createContextualFragment(e);return r.childNodes[0]}function ue(e){var r=h.createElement(\"body\");return r.innerHTML=e,r.childNodes[0]}function fe(e){return e=e.trim(),le?se(e):oe?de(e):ue(e)}function C(e,r){var t=e.nodeName,n=r.nodeName,a,i;return t===n?!0:(a=t.charCodeAt(0),i=n.charCodeAt(0),a<=90&&i>=97?t===n.toUpperCase():i<=90&&a>=97?n===t.toUpperCase():!1)}function ce(e,r){return!r||r===ie?h.createElement(e):h.createElementNS(r,e)}function ve(e,r){for(var t=e.firstChild;t;){var n=t.nextSibling;r.appendChild(t),t=n}return r}function G(e,r,t){e[t]!==r[t]&&(e[t]=r[t],e[t]?e.setAttribute(t,\"\"):e.removeAttribute(t))}var j={OPTION:function(e,r){var t=e.parentNode;if(t){var n=t.nodeName.toUpperCase();n===\"OPTGROUP\"&&(t=t.parentNode,n=t&&t.nodeName.toUpperCase()),n===\"SELECT\"&&!t.hasAttribute(\"multiple\")&&(e.hasAttribute(\"selected\")&&!r.selected&&(e.setAttribute(\"selected\",\"selected\"),e.removeAttribute(\"selected\")),t.selectedIndex=-1)}G(e,r,\"selected\")},INPUT:function(e,r){G(e,r,\"checked\"),G(e,r,\"disabled\"),e.value!==r.value&&(e.value=r.value),r.hasAttribute(\"value\")||e.removeAttribute(\"value\")},TEXTAREA:function(e,r){var t=r.value;e.value!==t&&(e.value=t);var n=e.firstChild;if(n){var a=n.nodeValue;if(a==t||!t&&a==e.placeholder)return;n.nodeValue=t}},SELECT:function(e,r){if(!r.hasAttribute(\"multiple\")){for(var t=-1,n=0,a=e.firstChild,i,d;a;)if(d=a.nodeName&&a.nodeName.toUpperCase(),d===\"OPTGROUP\")i=a,a=i.firstChild;else{if(d===\"OPTION\"){if(a.hasAttribute(\"selected\")){t=n;break}n++}a=a.nextSibling,!a&&i&&(a=i.nextSibling,i=null)}e.selectedIndex=t}}},E=1,pe=11,Y=3,Q=8;function N(){}function he(e){if(e)return e.getAttribute&&e.getAttribute(\"id\")||e.id}function ge(e){return function(t,n,a){if(a||(a={}),typeof n==\"string\")if(t.nodeName===\"#document\"||t.nodeName===\"HTML\"||t.nodeName===\"BODY\"){var i=n;n=h.createElement(\"html\"),n.innerHTML=i}else n=fe(n);var d=a.getNodeKey||he,v=a.onBeforeNodeAdded||N,f=a.onNodeAdded||N,b=a.onBeforeElUpdated||N,p=a.onElUpdated||N,y=a.onBeforeNodeDiscarded||N,A=a.onNodeDiscarded||N,K=a.onBeforeElChildrenUpdated||N,L=a.childrenOnly===!0,S=Object.create(null),H=[];function D(s){H.push(s)}function X(s,o){if(s.nodeType===E)for(var l=s.firstChild;l;){var u=void 0;o&&(u=d(l))?D(u):(A(l),l.firstChild&&X(l,o)),l=l.nextSibling}}function x(s,o,l){y(s)!==!1&&(o&&o.removeChild(s),A(s),X(s,l))}function q(s){if(s.nodeType===E||s.nodeType===pe)for(var o=s.firstChild;o;){var l=d(o);l&&(S[l]=o),q(o),o=o.nextSibling}}q(t);function V(s){f(s);for(var o=s.firstChild;o;){var l=o.nextSibling,u=d(o);if(u){var g=S[u];g&&C(o,g)?(o.parentNode.replaceChild(g,o),U(g,o)):V(o)}else V(o);o=l}}function te(s,o,l){for(;o;){var u=o.nextSibling;(l=d(o))?D(l):x(o,s,!0),o=u}}function U(s,o,l){var u=d(o);u&&delete S[u],!(!l&&(b(s,o)===!1||(e(s,o),p(s),K(s,o)===!1)))&&(s.nodeName!==\"TEXTAREA\"?re(s,o):j.TEXTAREA(s,o))}function re(s,o){var l=o.firstChild,u=s.firstChild,g,m,M,B,T;e:for(;l;){for(B=l.nextSibling,g=d(l);u;){if(M=u.nextSibling,l.isSameNode&&l.isSameNode(u)){l=B,u=M;continue e}m=d(u);var P=u.nodeType,w=void 0;if(P===l.nodeType&&(P===E?(g?g!==m&&((T=S[g])?M===T?w=!1:(s.insertBefore(T,u),m?D(m):x(u,s,!0),u=T):w=!1):m&&(w=!1),w=w!==!1&&C(u,l),w&&U(u,l)):(P===Y||P==Q)&&(w=!0,u.nodeValue!==l.nodeValue&&(u.nodeValue=l.nodeValue))),w){l=B,u=M;continue e}m?D(m):x(u,s,!0),u=M}if(g&&(T=S[g])&&C(T,l))s.appendChild(T),U(T,l);else{var k=v(l);k!==!1&&(k&&(l=k),l.actualize&&(l=l.actualize(s.ownerDocument||h)),s.appendChild(l),V(l))}l=B,u=M}te(s,u,m);var W=j[s.nodeName];W&&W(s,o)}var c=t,_=c.nodeType,J=n.nodeType;if(!L){if(_===E)J===E?C(t,n)||(A(t),c=ve(t,ce(n.nodeName,n.namespaceURI))):c=n;else if(_===Y||_===Q){if(J===_)return c.nodeValue!==n.nodeValue&&(c.nodeValue=n.nodeValue),c;c=n}}if(c===n)A(t);else{if(n.isSameNode&&n.isSameNode(c))return;if(U(c,n,L),H)for(var I=0,ne=H.length;I<ne;I++){var $=S[H[I]];$&&x($,$.parentNode,!1)}}return!L&&c!==t&&t.parentNode&&(c.actualize&&(c=c.actualize(t.ownerDocument||h)),t.parentNode.replaceChild(c,t)),c}}var be=ge(ae),Z=be;function F(e){let r=e.starter;if(e.id){let t=document.getElementById(e.id);if(!t)throw new Error(`Error while locating root with id: can't find direct with ${e}`);r=t}else{let t=0;for(;;){if(!r.parentElement)throw new Error(`Error while locating root: can't find parent with ${e}`);if(!r.getAttribute(\"state\"))r=r.parentElement;else if(e.depth&&t!=e.depth)r=r.parentElement,t++;else break}}return r}function Ae(e){return e.includes(\":\")&&(e=e.split(\":\")[1]),e.includes(\"$\")&&(e=e.replaceAll(\"$\",\"\")),e}function me(e){e.querySelectorAll(\"[ssa\\\\:oncall\\\\.display]\").forEach(t=>{let n=t.getAttribute(\"ssa:oncall.display\");n!=\"\"&&t.setAttribute(\"style\",\"display: \"+n)})}function Te(){document.querySelectorAll(\"[ssa\\\\:onload]\").forEach(e=>{let r=e.getAttribute(\"ssa:onload\");r&&r!=\"\"&&O(e,r)})}function we(){document.querySelectorAll(\"[ssa\\\\:poll]\").forEach(e=>{let r=e.getAttribute(\"ssa:poll\")||\"\",t=e.getAttribute(\"ssa:poll.interval\");r&&r!=\"\"&&t&&t!=\"\"&&setInterval(()=>{O(e,r)},parseInt(t))})}function ye(){document.querySelectorAll(\"[ssa\\\\:onintersect]\").forEach(e=>{let r=e.getAttribute(\"ssa:onintersect\")||\"\",t=e.getAttribute(\"ssa:onintersect.threshold\")||\"1.0\";r!=\"\"&&new IntersectionObserver(a=>{a.forEach(i=>{i.intersectionRatio>=parseFloat(t)&&O(e,r,parseFloat(t))})},{threshold:parseFloat(t)}).observe(e)})}function ee(e,r,t){let n=new Array,a={};t&&(a=t),a.onBeforeElUpdated=function(i,d){if(i.getAttribute(\"ssa:morph.ignore.attr\")!=null){let v=i.getAttribute(\"ssa:morph.ignore.attr\");if(v)if(v==\"innerHTML\")d.innerHTML=i.innerHTML;else{let f=i.getAttribute(v);f&&d.setAttribute(v,f)}}return i.getAttribute(\"ssa:morph.ignore\")!=null?!1:i.getAttribute(\"ssa:morph.ignore.this\")!=null&&i!=e?(n.push({fromEl:i,toEl:d}),!1):!0},Z(e,r,a),n.length>0&&n.forEach(i=>{ee(i.fromEl,i.toEl,{childrenOnly:!0})})}function O(e,r,...t){return new Promise((n,a)=>{let i=F({starter:e,depth:r.split(\"\").filter(y=>y===\"$\").length,id:r.includes(\":\")?r.split(\":\")[0]:void 0});me(i);let d=actionpath;d.endsWith(\"/\")||(d+=\"/\"),d+=`${i.getAttribute(\"name\")}`,d+=`/${Ae(r)}`;let v=new FormData;v.set(\"State\",i.getAttribute(\"state\")),v.set(\"Args\",JSON.stringify(t));var f=new XMLHttpRequest;f.open(\"POST\",d,!0);let b=0,p=\"\";f.onprogress=function(){let y=f.responseText.length;if(b==y)return;let A=this.responseText.substring(b,y);if(A.startsWith(\"ssa:redirect=\")){window.location.href=A.replace(\"ssa:redirect=\",\"\");return}if(p+=A,p.endsWith(actionterminator)){switch(p=p.slice(0,-actionterminator.length),i.getAttribute(\"ssa:render.mode\")||\"morph\"){case\"replace\":i.outerHTML=p;break;case\"morph\":try{ee(i,p)}catch(L){console.log('Fallback from \"morphdom\" to \"replace\" due to an error:',L),i.outerHTML=p}break;default:console.log('Render mode is not supported, fallback to \"replace\"'),i.outerHTML=p;break}p=\"\"}b=y},f.onload=function(){n()},f.onerror=function(){a()},f.onabort=function(){a()},f.send(v)})}function Ne(e,r){let t=F({starter:e,depth:r.split(\"\").filter(a=>a===\"$\").length,id:r.includes(\":\")?r.split(\":\")[0]:void 0});if(!t.getAttribute(\"state\"))throw new Error(\"Bind call error: component state is underfined\");let n=JSON.parse(atob(t.getAttribute(\"state\")));n[r]=e.value,t.setAttribute(\"state\",btoa(JSON.stringify(n)))}function Se(e,r){r.preventDefault();let t=F({starter:e});if(!t.getAttribute(\"state\"))throw new Error(\"Bind call error: component state is underfined\");let n=JSON.parse(atob(t.getAttribute(\"state\"))),a=new FormData(r.target),i=Object.fromEntries(a.entries());return Object.entries(i).forEach(d=>{n[d[0]]=d[1]}),t.setAttribute(\"state\",btoa(JSON.stringify(n))),O(t,\"Submit\"),!1}window._root=F;window.Action=O;window.Bind=Ne;window.FormSubmit=Se;document.addEventListener(\"DOMContentLoaded\",Te);document.addEventListener(\"DOMContentLoaded\",ye);document.addEventListener(\"DOMContentLoaded\",we);})();</script>"
