<script>
  import { beforeUpdate } from 'svelte';
  import {grpc} from "@improbable-eng/grpc-web";
  import {ControlMsgEmpty} from '../../pb/control_pb';
  import {ControlMsgServiceClient} from "../../pb/control_pb_service";
  import { controlCli, date, cpuUsage, memUsage } from '$lib/stores'

  // grpc.setDefaultTransport(grpc.CrossBrowserHttpTransport({withCredentials: false}));
  const doRefresh = () => $controlCli.refresh(new ControlMsgEmpty())

  beforeUpdate(() => {
      if ($controlCli.control === undefined) {
          console.log("beforeUp: ", $controlCli)
          controlCli.set(new ControlMsgServiceClient(window.location.origin, {
              transport: grpc.WebsocketTransport()
          }))

          $controlCli.control(new ControlMsgEmpty())
              .on('data', (cmsg) => {
                  let obj = cmsg.toObject()
                  date.set(obj.date)
                  cpuUsage.set(obj.cpuUsage)
                  memUsage.set(obj.memUsage)
              })
              .on('status', (x) => console.log("on Status: ", x))
              .on('end', (x) => console.log("on End: ", x))

          console.log("ControlMsg Client", $controlCli)
      }
  })

</script>

{#if $date != ""}
  <div class="block">
    <p>Streaming from server .... {$date} -- cpu: {$cpuUsage} | mem: {$memUsage}</p>
  </div>
{/if}
<button class="button" on:click={doRefresh}>Refresh</button>
