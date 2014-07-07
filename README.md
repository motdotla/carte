# carte

<img src="https://raw.githubusercontent.com/scottmotte/carte/master/carte.png" alt="carte" width="190" />

API of memory cards. I'm beginning this with people in my company.

## Usage

```
git clone https://github.com/scottmotte/carte.git
cd carte
```

Then run [srvdir](https://srvdir.net/) and visit the url.

```
srvdir
```

## Data Scraping

### From Parklet

[Parklet](http://parklet.co) might be where you check your employee directory. Here's how to scrape all their photos and names with js.

Visit the [directory](https://app.parklet.co/directory), go to the view that has the smaller pictures, load everyone, and then open Chrome console.

Paste the following in.

```javascript
var results = []; var uri_pattern = /\b((?:[a-z][\w-]+:(?:\/{1,3}|[a-z0-9%])|www\d{0,3}[.]|[a-z0-9.\-]+[.][a-z]{2,4}\/)(?:[^\s()<>]+|\(([^\s()<>]+|(\([^\s()<>]+\)))*\))+(?:\(([^\s()<>]+|(\([^\s()<>]+\)))*\)|[^\s`!()\[\]{};:'".,<>?«»“”‘’]))/ig; $item = $(".value-items .value-item"); $item.each(function() { var style = $( this ).find(".employee").attr("style"); var matches = style.match(uri_pattern); var name = $(this).find("span[name='name']").text(); if (matches) { results.push({front: "<img src='"+matches[0]+"' />", back: name, tags: ["sendgrid"] });} }); var str = JSON.stringify(results, undefined, 2); console.log(str); $("body").append(str);
```

Scroll to the bottom of the screen and double click the newly appended content. 

Paste that json content into [here](http://jsonformat.com/).

Copy the results. Paste those results in cards.json.

## API Overview

The [carte.herokuapp.com](https://carte.herokuapp.com) API is based around REST. It uses standard HTTP authentication. JSON is returned in all responses from the API, including errors.

I've tried to make it as easy to use as possible, but if you have any feedback please [let me know](mailto:scott@scottmotte.com).

## Summary

### API Endpoint

* [https://carte-api.herokuapp.com/api/v0](https://carte-api.herokuapp.com/api/v0)

### Accounts

To start using the carte API, you must first create a deck.

#### ANY /accounts/create 

Pass an email to create your account at carte-api.herokuapp.com.

##### Definition

```
ANY https://carte-api.herokuapp.com/api/v0/accounts/create.json?&email=[email]&api_key=[api_key]
```

##### Required Parameters

* email

##### Optional Parameters

* api_key

##### Example Request

<https://carte-api.herokuapp.com/api/v0/accounts/create.json?email=[email]&api_key=[api_key]>

##### Example Response

```
{
  "accounts": [{
    "email": "test@example.com",
    "api_key": "the_default_generated_api_key_that_you_should_keep_secret"
  }]
}
```

##### Example Error

```
{
  errors: [{
    "code": "required",
    "field": "email",
    "message": "email cannot be blank"
  }]
}
```

#### ANY /cards/create

Pass an api_key, front, and back to create your card.

##### Definition

```
ANY https://carte-api.herokuapp.com/api/v0/cards/create.json?api_key=[api_key]&front=[front]&back=[back]
```

##### Required Parameters

* api_key
* front
* back

##### Example Request

<https://carte-api.herokuapp.com/api/v0/cards/create.json?api_key=[api_key]&front=[front]&back=[back]>

##### Example Response

```
{
  "cards": [{
    "id": "12345",
    "front": "<img src='http://example.com/some-image.jpg'>",
    "back": "John Doe",
  }]
}
```

##### Example Error

```
{
  errors: [{
    "code": "required",
    "field": "front",
    "message": "front cannot be blank"
  }]
}
```

