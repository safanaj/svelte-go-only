<script>
  import {grpc} from "@improbable-eng/grpc-web";
  import {Categories, IndexRequest, IndexRequestKind} from '../../pb/categories_pb'
  import {CategoryService , CategoryServiceClient} from "../../pb/categories_pb_service";
  import { clientSide, serverSide } from '$lib/stores'

  import { Col, Label, FormGroup, Input, Button } from 'sveltestrap';

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

<Col class="jcc">
  <Label><p>Categories kind <Input inline type="switch" label="(auto change)" bind:checked={autoChangeKind} /></p>
    <Col>
      <label for="serverSide" class="form-label">Number of entries from server-side</label>
      <input type="range" class="form-range" id="serverSide" min="1" max="10000"
             bind:value={srvSide}
             on:input={() => serverSide.set(srvSide)} />
      <p>{$serverSide}</p>
    </Col>
    <Col>
      <label for="serverSide" class="form-label">Number of entries on client-side</label>
      <input type="range" class="form-range" id="clientSide" min="1" max="1000"
             bind:value={cliSide}
             on:input={() => clientSide.set(cliSide)} />
      <p>{$clientSide}</p>
    </Col>

    <FormGroup>
      {#each Object.entries(IndexRequestKind) as e}
        <Input type="radio" value={e[1]} label="{e[0].capitalize()}" bind:group={lastKind} />
      {/each}
    </FormGroup>
  </Label>
</Col>

<Col>
  <Label>
    <button on:click={getCategories}>Get random collection of a category (RPC Service Client)</button>
    <ul>
      {#each cats.slice(0, $clientSide) as cat}
	    <li>
	      {cat.id} - {cat.name}
	    </li>
      {/each}
    </ul>
  </Label>
</Col>
