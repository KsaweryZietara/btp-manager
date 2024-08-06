import * as ui5 from "@ui5/webcomponents-react";
import { ApiError, ServiceInstanceBindings } from "../shared/models";
import axios from "axios";
import { useEffect, useState } from "react";
import api from "../shared/api";
import Ok from "../shared/validator";
import serviceInstancesData from '../test-data/service-bindings.json';
import StatusMessage from "./StatusMessage";

function ServiceBindingsList(props: any) {
  const [bindings, setServiceInstanceBindings] = useState<ServiceInstanceBindings>();

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<ApiError>();
  const [portal] = useState<JSX.Element>();
  const [success] = useState("");

  function deleteBinding(id: string): boolean {
    setLoading(true);

    axios
      .delete(api("service-bindings") + "/" + id)
      .then((response) => {
        setLoading(false);
      })
      .catch((error) => {
        setLoading(false);
        setError(error);
      });

    return true;
  }



  useEffect(() => {
    if (!Ok(props.instance)) {
      return;
    }

    var useTestData = process.env.REACT_APP_USE_TEST_DATA === "true"
    if (!useTestData) {
      setLoading(true)
      axios
        .get<ServiceInstanceBindings>(api("service-bindings"),
          { params: { service_instance_id: props.instance.id } }
        )
        .then((response) => {
          setServiceInstanceBindings(response.data);
          setLoading(false);
        })
        .catch((error) => {
          setLoading(false);
          setError(error);
        });
      setLoading(false)
    } else {
      setLoading(true)
      setServiceInstanceBindings(serviceInstancesData)
      setLoading(false);
    }
  }, [props.instance]);

  if (loading) {
    return <ui5.BusyIndicator
      active
      delay={1000}
      size="Medium"
    />
  }

  const renderData = () => {
    // @ts-ignore
    if (!Ok(bindings) || !Ok(bindings.items)) {
      return <ui5.IllustratedMessage name="NoEntries" />
    }
    return bindings?.items.map((binding, index) => {
      return (
        <>
          <ui5.TableRow>

            <ui5.TableCell>
              <ui5.Label>{binding.id}</ui5.Label>
            </ui5.TableCell>

            <ui5.TableCell>
              <ui5.Label>{binding.name}</ui5.Label>
            </ui5.TableCell>

            <ui5.TableCell>
              <ui5.Label>{binding.namespace}</ui5.Label>
            </ui5.TableCell>

            <ui5.TableCell>
              <ui5.Button
                design="Default"
                icon="delete"
                onClick={function _a(e: any) {
                  e.stopPropagation();
                  return deleteBinding(binding.id);
                }}
              >
              </ui5.Button>            
            </ui5.TableCell>

          </ui5.TableRow>
        </>
      );
    });
  };

  return (
    <>
      <ui5.Form>
      <StatusMessage error={error ?? undefined} success={success} />
</ui5.Form>

      {

        <ui5.Table
          columns={
            <>
              <ui5.TableColumn>
                <ui5.Label>Id</ui5.Label>
              </ui5.TableColumn>

              <ui5.TableColumn>
                <ui5.Label>Name</ui5.Label>
              </ui5.TableColumn>

              <ui5.TableColumn>
                <ui5.Label>Namespace</ui5.Label>
              </ui5.TableColumn>

              <ui5.TableColumn>
                <ui5.Label>Action</ui5.Label>
              </ui5.TableColumn>
            </>
          }
        >
          {renderData()}
        </ui5.Table>
      }

      {portal != null && portal}

    </>
  );
}

export default ServiceBindingsList;