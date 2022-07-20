<script>
  import {grpc} from "@improbable-eng/grpc-web";
  import {Categories, IndexRequest, IndexRequestKind} from '../../pb/categories_pb'
  import {CategoryService , CategoryServiceClient} from "../../pb/categories_pb_service";

  import { Col, Label, FormGroup, Input, Button } from 'sveltestrap';

  let cats = [];
  let lastKind = IndexRequestKind.CITY;
  let autoChangeKind = false;
  let cli = new CategoryServiceClient(window.location.protocol + '//' + window.location.host);

  const getCategories = async () => {
      let indexReq = new IndexRequest()
      indexReq.setKind(lastKind)
      console.log(`Kind Set to: ${lastKind} , is auto change enabled: ${autoChangeKind}`)
      if (autoChangeKind) {
        lastKind = (lastKind + 1) % Object.keys(IndexRequestKind).length
        // if (lastKind == IndexRequestKind.CITY) {
        //   lastKind = IndexRequestKind.COUNTRY
        // } else {
        //   lastKind = IndexRequestKind.CITY
        // }
      }
      cli.index(indexReq, (err, res) => {
          if (err === null) {
              cats = res.toObject().categoriesList
          } else { console.log(err) }
      })
  }

</script>

<Col>
<Label><p>Categories kind <Input inline type="switch" label="(auto change)" bind:checked={autoChangeKind} /></p>
  <FormGroup>
    {#each Object.entries(IndexRequestKind) as e}
      <Input type="radio" value={e[1]} label="{e[0].capitalize()}" bind:group={lastKind} />
    {/each}
  </FormGroup>
</Label>
</Col>

<Col>
<button on:click={getCategories}>Get random cities/countries (RPC Service Client)</button>

<ul>
  {#each cats.slice(0,5) as cat}
	<li>
	  {cat.id} - {cat.name}
	</li>
  {/each}
</ul>
</Col>
