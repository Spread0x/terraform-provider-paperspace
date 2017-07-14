package main

import (
  "encoding/json"
  "fmt"
  "github.com/hashicorp/terraform/helper/schema"
  "log"
)

func dataSourceTemplateRead(d *schema.ResourceData, m interface{}) error {
  client := m.(PaperspaceClient).RestyClient

  log.Printf("[INFO] paperspace dataSourceTemplateRead Client ready")

  queryParam := false;
  queryStr := "?"
  id, ok := d.GetOk("id")
  if ok {
    queryStr += "id=" + id.(string)
    queryParam = true
  }
  name, ok := d.GetOk("name")
  if ok {
    if queryParam {
      queryStr += "&"
    }
    queryStr += "name=" + name.(string)
    queryParam = true
  }
  label, ok := d.GetOk("label")
  if ok {
    if queryParam {
      queryStr += "&"
    }
    queryStr += "label=" + label.(string)
    queryParam = true
  }
  os, ok := d.GetOk("os")
  if ok {
    if queryParam {
      queryStr += "&"
    }
    queryStr += "os=" + os.(string)
    queryParam = true
  }
  dtCreated, ok := d.GetOk("dtCreated")
  if ok {
    if queryParam {
      queryStr += "&"
    }
    queryStr += "dtCreated=" + dtCreated.(string)
    queryParam = true
  }
  teamId, ok := d.GetOk("teamId")
  if ok {
    if queryParam {
      queryStr += "&"
    }
    queryStr += "teamId=" + teamId.(string)
    queryParam = true
  }
  userId, ok := d.GetOk("userId")
  if ok {
    if queryParam {
      queryStr += "&"
    }
    queryStr += "userId=" + userId.(string)
    queryParam = true
  }
  region, ok := d.GetOk("region")
  if ok {
    if queryParam {
      queryStr += "&"
    }
    queryStr += "region=" + region.(string)
    queryParam = true
  }
  if !queryParam {
    return fmt.Errorf("Error reading paperspace template: must specify query filter properties")
  }

  resp, err := client.R().
  Get("/templates/getTemplates" + queryStr)
  if err != nil {
    return fmt.Errorf("Error reading paperspace template: %s", err)
  }

  statusCode := resp.StatusCode()
  log.Printf("[INFO] paperspace dataSourceTemplateRead StatusCode: %v", statusCode)
  LogResponse("paperspace dataSourceTemplateRead", resp, err)
  if statusCode == 404 {
    return fmt.Errorf("Error reading paperspace template: templates not found")
  }
  if statusCode != 200 {
    return fmt.Errorf("Error reading paperspace template: Response: %s", resp.Body())
  }

  var f interface{}
  err = json.Unmarshal(resp.Body(), &f)
  if err != nil {
    return fmt.Errorf("Error unmarshalling paperspace template read response: %s", err)
  }

  mpa := f.([]interface{})
  if len(mpa) > 1 {
    return fmt.Errorf("Error reading paperspace template: found more than one template matching given properties")
  }
  if len(mpa) == 0 {
    return fmt.Errorf("Error reading paperspace template: no template found matching given properties")
  }

  mp, ok := mpa[0].(map[string]interface{})
  if !ok {
    return fmt.Errorf("Error unmarshalling paperspace template read response: no templates not found")
  }

  idr, _ := mp["id"].(string)
  if idr == "" {
    return fmt.Errorf("Error unmarshalling paperspace template read response: no template id found for template")
  }

  log.Printf("[INFO] paperspace dataSourceTemplateRead template id: %v", idr)

  SetResData(d, mp, "name")
  SetResData(d, mp, "label")
  SetResData(d, mp, "os")
  SetResData(d, mp, "dtCreated")
  SetResData(d, mp, "teamId")
  SetResData(d, mp, "userId")
  SetResData(d, mp, "region")

  d.SetId(idr)

	return nil
}

func dataSourceTemplate() *schema.Resource {
  return &schema.Resource{
    Read: dataSourceTemplateRead,

		Schema: map[string]*schema.Schema{
      "id": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "name": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "label": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "os": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "dtCreated": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "teamId": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "userId": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "region": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
    },
	}
}
