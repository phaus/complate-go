import { Card } from "./components.jsx";
import { createElement } from "complate-stream";

export function SampleView() {
	return <Card title="Hello World">
		<p>lorem ipsum dolor sit amet</p>
	</Card>;
}