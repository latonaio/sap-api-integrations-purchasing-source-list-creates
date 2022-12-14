# sap-api-integrations-purchasing-source-list-creates  
sap-api-integrations-purchase-contract-creates は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 供給元一覧データ を登録するマイクロサービスです。  
sap-api-integrations-purchase-contract-creates には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-purchase-contract-creates は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_PURCHASING_SOURCE_LIST/overview  

## 動作環境  
sap-api-integrations-purchasing-source-list-creates は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）   
・ AION のリソース （推奨)   
・ OS: LinuxOS （必須）   
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用
sap-api-integrations-purchasing-source-list-creates は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。


## 本レポジトリ が 対応する API サービス
sap-api-integrations-purchasing-source-list-creates が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PURCHASING_SOURCE_LIST_SRV/overview  
* APIサービス名(=baseURL): API_PURCHASING_SOURCE_LIST_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-purchasing-source-list-creates には、次の API をコールするためのリソースが含まれています。  

* A_PurchasingSourceList（供給元一覧）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"PurchasingSourceList" が指定されています。    
  
```
"api_schema": "SAPPurchasingSourceListCreates",
"accepter": ["PurchasingSourceList"],
"purchasing_source_list": "",
"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
"api_schema": "SAPPurchasingSourceListCreates",
"accepter": ["All"],
"purchasing_source_list": "",
"deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostList(
	list *requests.PurchasingSourceList,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "List":
			func() {
				c.List(purchasingSourceList)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 供給元一覧 が登録された結果の JSON の例です。  
以下の項目のうち、"Material" ～ "SourceOfSupplyIsFixed" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-purchasing-source-list-creates/SAP_API_Caller/caller.go#L205",
	"function": "sap-api-integrations-purchasing-source-list-creates/SAP_API_Caller.(*SAPAPICaller).PurchasingSourceList",
	"level": "INFO",
	"message": [
		{
			"Material": "RM128",
			"Plant": "1710",
			"SourceListRecord": "1",
			"ValidityStartDate": "2016-07-01T09:00:00+09:00",
			"ValidityEndDate": "2020-06-21T09:00:00+09:00",
			"Supplier": "17300001",
			"PurchasingOrganization": "1710",
			"SupplyingPlant": "",
			"OrderQuantityUnit": "PC",
			"PurchaseOutlineAgreement": "5600001710",
			"SupplierIsFixed": false,
			"SourceOfSupplyIsBlocked": false,
			"MRPSourcingControl": "2",
			"LastChangeDateTime": "",
			"IssgPlantIsFixed": false,
			"PurOutlineAgreementIsFixed": true,
			"SourceOfSupplyIsFixed": true
		}
	],
	"time": "2022-09-09T10:59:37.566238+09:00"
}
```
