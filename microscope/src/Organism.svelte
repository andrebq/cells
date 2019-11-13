<script context="module">
  let subsmap = new Map();

  export function subscribe(cellid, callback) {
    let subscription = { id: cellid, callback: callback };
    if (subsmap.has(cellid)) {
      subsmap.get(cellid).push(subscription);
      return;
    }

    subsmap.set(cellid, [subscription]);
  }

  export function unsubscribe(subscription) {
    if (!subsmap.has(subscription.id)) {
      return;
    }
    items = subsmap.get(subscription.id);
    items.remove(subscribe);
  }

  export function update(cellid, value) {
    if (!subsmap.has(cellid)) {
      return;
    }
    subsmap
      .get(cellid)
      .forEach((subscription, _idx) => subscription.callback(value));
  }
</script>
<section></section>
