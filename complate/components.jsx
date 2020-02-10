import { createElement } from "complate-stream";

export function Card({ title }, ...children) {
	return <article class="card">
		<h3>{title}</h3>
		{children}
	</article>;
}