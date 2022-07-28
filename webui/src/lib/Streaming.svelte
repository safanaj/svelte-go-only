<script>
  import { beforeUpdate } from 'svelte';
  import {grpc} from "@improbable-eng/grpc-web";
  import {ControlMsgEmpty} from '../../pb/control_pb';
  import {ControlMsgServiceClient} from "../../pb/control_pb_service";

  let date = "", cpuU = "", memU = ""
  let controlCli = null

  // grpc.setDefaultTransport(grpc.CrossBrowserHttpTransport({withCredentials: false}));
  const doRefresh = () => controlCli.refresh(new ControlMsgEmpty())

  beforeUpdate(() => {
      if (controlCli === null) {
          controlCli = new ControlMsgServiceClient(window.location.origin, {
              transport: grpc.WebsocketTransport()
          })
          controlCli.control(new ControlMsgEmpty())
              .on('data', (cmsg) => {
                  console.log("on data: ", cmsg)
                  let obj = cmsg.toObject()
                  date = obj.date
                  cpuU = obj.cpuUsage
                  memU = obj.memUsage
              })
              .on('status', (x) => console.log("on Status: ", x))
              .on('end', (x) => console.log("on End: ", x))

          console.log("ControlMsg Client", controlCli)
      }
  })

</script>
{#if date != ""}
  <p>Streaming from server .... {date} -- cpu: {cpuU} | mem: {memU}</p>
{/if}
<button on:click={doRefresh}>Refresh</button>
