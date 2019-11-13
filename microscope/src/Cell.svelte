<script>
    import Organism, {subscribe, unsubscribe} from './Organism.svelte';
    export let size = "m";
    export let title = "";
    export let cellid = "";
    import { onMount } from 'svelte';

    let subscription = null;
    let value = null;

    onMount(() => {
        if (cellid) {
            console.info("cellid: ", cellid);
            subscribe(cellid, (new_value) => {
                console.info("value: ", value, "new value", new_value);
                value = new_value;
            });
        }
    });
</script>

<section class="cell-tile {size}">
    <section class="cell">
        <h1 class="title">{title}</h1>
        <p>{value ? value.key : ''}</p>
    </section>
</section>

<style>
    .title {
        color: #556e53;
        margin: 0;
    }

    p {
        margin: 0;
        padding: 0;
    }

    .cell {
        width: 100%;
        height: 100%;
        background-color: #1e3544;
        overflow: hidden;
    }

    .cell:hover {
        background-color: #29435c;
    }

    .cell-tile {
        display: flex;
        width: 100px;
        height: 100px;
        padding: 10px;
        margin: 0;
    }

    .cell-tile.xl{
        width: 200px;
        height: 200px;
    }
    .cell-tile.xl.wide {
        width: 400px;
        height: 200px;
    }
    .cell-tile.l {
        width: 150px;
        height: 150px;
    }
    .cell-tile.xs {
        width: 50px;
        height: 50px;;
    }
    .cell-tile.xs.wide {
        width: 100px;
        height: 50px;;
    }
    .cell-tile.xs .title {
        display: none;
    }
</style>
