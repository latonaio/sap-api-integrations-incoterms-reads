# sap-api-integrations-incoterms-reads
sap-api-integrations-incoterms-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で インコタームズデータを取得するマイクロサービスです。    
sap-api-integrations-incoterms-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-incoterms-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_SD_INCOTERMS_SRV_0001/overview  

## 動作環境  
sap-api-integrations-incoterms-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-incoterms-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-incoterms-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SD_INCOTERMS_SRV_0001/overview    
* APIサービス名(=baseURL): API_SD_INCOTERMS_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-incoterms-reads には、次の API をコールするためのリソースが含まれています。  

* A_IncotermsClassification（インコタームズ - インコタームズ分類）※インコタームズ分類の詳細データを取得するために、ToIncotermsClassificationText、と合わせて利用されます。
* ToIncotermsClassificationText（インコタームズ - インコタームズ分類テキスト ※To）
* A_IncotermsVersion（インコタームズ - インコタームズバージョン）※インコタームズバージョンの詳細データを取得するために、ToIncotermsVersionText、と合わせて利用されます。
* ToIncotermsVersionText（インコタームズ - インコタームズバージョンテキスト ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-incoterms-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inputSDC.IncotermsClassification.IncotermsClassification,
* inputSDC.IncotermsVersion.IncotermsVersion,

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"IncotermsClassification" が指定されています。

```
	"api_schema": "SAPIncotermsReads",
	"accepter": ["IncotermsClassification"],
	"incoterms_code": "CFR",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPIncotermsReads",
	"accepter": ["All"],
	"incoterms_code": "CFR",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetIncoterms(incotermsClassification, incotermsVersion string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "IncotermsClassification":
			func() {
				c.IncotermsClassification(incotermsClassification)
				wg.Done()
			}()
		case "IncotermsVersion":
			func() {
				c.IncotermsVersion(incotermsVersion)
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
以下の sample.json の例は、SAP インコタームズ の インコタームズ分類データ が取得された結果の JSON の例です。  
以下の項目のうち、"IncotermsClassification" ～ "to_IncotermsClassificationText" は、/SAP_API_Output_Formatter/type.go 内 の Type IncotermsClassification {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
XXX
```
