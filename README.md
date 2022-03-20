# goForMercaing

### A mercari V1 API wrapper rewritten in golang(same as [marvinody/mercari](https://github.com/marvinody/mercari))

## Usage:  
`    mercarigo.Mercari_search(name string, sort string, order string, status string, limit int, times int) ([]MercariItem, error)`

## Params:  
| Param  | Description|
| ---    | ---        |
| name   | item to search.  |
| sort   | there would be other options, but created time is enough.  
| order  | using desc to get latest info, which is enough.  |
| status | item status ( on_sale / trading / sold_out / leave empty string )  |
| limit  | items limit on one page, collaborate with times to control total item count. |  
| times  | maximum page count when searching, collbrate with limit to control total item count. |  

## Return :  
`[]MercariItem`  
|Variables | Description|
| --- | --- |
|ProductId   | a unique identifier for results, you could get link by using `MercariItem.GetProductURL()`.  |
| ProductName | product name  |
| Price       | a int number, price in yen.|  
| Created     | unix time about create time.  |
| Updated     | unix time about item update.  |
| Condition[[*](#condition)]| a `Name_Id_Unit`, parameter id refers to item condition  |
| ImageURL    | first picture of item from mercdn  |
| Status      | info about item selling status( on_sale / trading / sold_out )  |
| Seller      | a `Name_Id_Unit`, parameter id is seller's id  |
| Buyer       | a `Name_Id_Unit`, similar to seller, only appear when item is not on sale.  |
| Shipping    | a `Name_Id_Unit`, id represents place shipping from, name is filled with usable string.  |

#### Condition explained in japanese
|Condition | Description|
|--- | ---|
|1 | 新品、未使用 | 
|2 | 未使用に近い |
|3 | 目立った傷や汚れなし  |
|4 | やや傷や汚れあり  |
|5 | 傷や汚れあり  |
|6 | 全体的に状態が悪い  |

error : errors during searching, including getting dPoP, sending and receiving http request, gzip decompression, json parse and more.  