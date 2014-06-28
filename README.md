# carte

Under develompent.

API of memory cards.

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
