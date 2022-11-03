<script>
  import {grpc} from "@improbable-eng/grpc-web";
  import {Categories, IndexRequest, IndexRequestKind} from '../../pb/categories_pb'
  import {CategoryService , CategoryServiceClient} from "../../pb/categories_pb_service";
  import { clientSide, serverSide } from '$lib/stores'

  let cats = [];
  let lastKind = 0;
  let autoChangeKind = false;
  let srvSide = $serverSide
  let cliSide = $clientSide
  let cli = new CategoryServiceClient(window.location.protocol + '//' + window.location.host);

  const getCategories = async () => {
      let indexReq = new IndexRequest()
      indexReq.setKind(lastKind)
      indexReq.setNumber($serverSide)
      if (autoChangeKind) {
        lastKind = (lastKind + 1) % Object.keys(IndexRequestKind).length
      }
      cli.index(indexReq, (err, res) => {
          if (err === null) {
              cats = res.toObject().categoriesList
          } else { console.log(err) }
      })
  }
</script>


<div class="block">
  <!-- <p>Automatically change categories kind <input type="checkbox" bind:checked={autoChangeKind} /></p> -->
  <!-- <label class="checkbox"> -->
  <!--   Automatically change categories kind -->
  <!--   <input class="switch" type="checkbox" bind:checked={autoChangeKind} > -->
  <!-- </label> -->
  <div class="field">
    <input id="autoChangeSwitch" type="checkbox" class="switch is-rtl is-info is-rounded" bind:checked={autoChangeKind} />
    <label for="autoChangeSwitch">Automatically change categories kind</label>
  </div>
</div>
<div class="columns">
  <div class="column">
    <div class="block">
      <label for="serverSide" class="form-label">Number of entries from server-side</label>
    </div>
    <div class="block">
      <input type="range" class="form-range" id="serverSide" min="1" max="10000"
             bind:value={srvSide}
             on:input={() => serverSide.set(srvSide)} />
      <p>{$serverSide}</p>
    </div>
  </div>

  <div class="column">
    <div class="block">
      <label for="clientSide" class="form-label">Number of entries on client-side</label>
    </div>
    <div class="block">
      <input type="range" class="form-range" id="clientSide" min="1" max="1000"
             bind:value={cliSide}
             on:input={() => clientSide.set(cliSide)} />
      <p>{$clientSide}</p>
    </div>
  </div>
</div>

<div class="columns">
  <div class="column is-3 is-offset-4">
    {#each Object.entries(IndexRequestKind) as e}
      <div class="columns">
        <div class="column">
          <div class="control">
            <input id="radio-{e[0]}" type="radio" value={e[1]} bind:group={lastKind} />
          </div>
        </div>
        <div class="column level">
          <div class="level-left">
            <label class="radio" for="radio-{e[0]}">{e[0].capitalize()}</label>
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>

<div class="container">
  <button class="button" on:click={getCategories}>Get random collection of a category (RPC Service Client)</button>
  <ul>
    {#each cats.slice(0, $clientSide) as cat}
	  <li>
	    {cat.id} - {cat.name}
	  </li>
    {/each}
  </ul>
</div>
