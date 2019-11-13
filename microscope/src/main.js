import App from './App.svelte';
import Organism, {subscribe, update} from './Organism.svelte';

console.info([Organism, subscribe, update]);

const app = new App({
	target: document.body,
	props: {
		name: 'world'
	}
});
document.body.style.padding = "0";
document.body.style.margin = "0";

subscribe('console', function(value) { console.info(value); })
update('xl', { 'key': 'value' });

window.updateFn = update;

setInterval(() => {
	update('xl', {'key': (new Date()).toISOString() })
}, 500)

export default app;
