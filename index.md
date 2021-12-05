
# Loan



## Indices

* [Ungrouped](#ungrouped)

  * [Add Loan](#1-add-loan)
  * [Approval Loan](#2-approval-loan)
  * [Get All Loan](#3-get-all-loan)
  * [Update Loan](#4-update-loan)


--------


## Ungrouped



### 1. Add Loan



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: localhost:7070/v1/loan/add
```



***Body:***

```js        
{
    "amount": 10
}
```



### 2. Approval Loan



***Endpoint:***

```bash
Method: PUT
Type: RAW
URL: localhost:7070/v1/loan/approval
```



***Body:***

```js        
{
    "id": 2,
    "state": "success",
    "version": 2
}
```



### 3. Get All Loan



***Endpoint:***

```bash
Method: GET
Type: RAW
URL: localhost:7070/v1/loan
```



### 4. Update Loan



***Endpoint:***

```bash
Method: PUT
Type: RAW
URL: localhost:7070/v1/loan/update
```



***Body:***

```js        
{
    "id": 2,
    "amount": 20,
    "version": 1
}
```



---
[Back to top](#loan)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2021-12-05 15:28:52 by [docgen](https://github.com/thedevsaddam/docgen)
