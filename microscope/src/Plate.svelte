<script>
    import Cell from './Cell.svelte';
</script>

<article class="plate">
    <Cell size="xl wide" cellid="xl" title="Title"></Cell>
    <Cell size="l"></Cell>
    <Cell size="xs wide" cellid="xl" title="another cell / same cell id"></Cell>
</article>

<style>
.plate {
    display: flex;
}
</style>
