# pack-calculator

## Description

This is an app to calculate the number of packs required to fulfill customer orders for a product that has different pack sizes. The challenge includes the following rules:

1) Only whole packs can be sent; packs cannot be broken open.
2) Send out no more items than necessary to fulfill the order.
3) Send out as few packs as possible to fulfill each order.

## Request  

```bash
curl http://localhost:8080/packs?items=501

```

## UI

After running the app, you can access the UI at http://localhost:8080

## Deployment
The application has been deployed to Vercel for testing and can be accessed at:

UI: https://packs-calculator-9jd076656-saras-projects-22a48c8d.vercel.app/

API: https://packs-calculator-9jd076656-saras-projects-22a48c8d.vercel.app/api/packs?items=501