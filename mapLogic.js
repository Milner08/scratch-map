var mapData = [];
var countryList = [];
var userID = '59e891525770861770a47a7d';

$(document).ready(function(){
  //Load the map SVG
  $('#svgContainer').load('media/world.svg', function(){
      scrapeCountiesFromMap();
      getVistedCountriesAndUpadateMap();

      $("path.land").click(function(){
        console.log("Path clicked");
        toggleCountry($(this));
      });
  });

  $('#scratchButton').click(function(){
    toggleCountry($("path[title="+$('#countrySearch').val()+"]"));
  });

  $('#countrySearch').change(function(){
    updateButton();
  });

  $('#countrySearch').typeahead({ source: countryList});
});

function updateButton(){
  if($.inArray($("#countrySearch").val(), mapData.visited_countries) !== -1){
    $("#scratchButton").text('Unscratch');
    $("#scratchButton").prop('disabled', false);
  }else if($.inArray($("#countrySearch").val(), countryList) !== -1){
    $("#scratchButton").text('Scratch');
    $("#scratchButton").prop('disabled', false);
  }else{
    $("#scratchButton").prop('disabled', true);
  }
}

function toggleCountry(path){
  console.log("Looking to toggle " + path );
  var country = $(path).attr('title');
  if($.inArray(country, mapData.visited_countries) !== -1){
    $(path).css({ fill: "#CCCCCC" });
    mapData.visited_countries = $.grep(mapData.visited_countries, function(value) {
      return value != country;
    });
  }else{
    $(path).css({ fill: "#4286f4" });
    mapData.visited_countries.push(country);
  }

  sendUpdatedMapData();
  updateButton();
  populateVistedCountriesTable();

}

function populateVistedCountriesTable(){
  mapData.visited_countries.sort();
  $('#vistedCountryList').empty();
  $('#vistedCountryList').append('<h2>Visted Countries</h2>');
  mapData.visited_countries.forEach(function(country){
    $('#vistedCountryList').append('<div class="country">'+country+'</div>');
  });
}

function sendUpdatedMapData(){
  console.log(JSON.stringify(mapData))
  $.ajax({
    method: "PUT",
    url: "http://localhost:8000/mapdata",
    data: JSON.stringify(mapData)
  }).done(function( msg ) {
    console.log("Updated successfully.")
  });
}

function scrapeCountiesFromMap(){
  $('#svgContainer').find('path.land').each(function(){
    countryList.push($(this).attr('title'));
  });

  countryList.sort();

  countryList.forEach(function(country){
    $('#countries').append($('<option>', {
      value: country,
      text: country
    }));
  })
}

function getVistedCountriesAndUpadateMap(){
  $.get("http://localhost:8000/mapdata/"+userID, function(data){
    mapData = $.parseJSON(data);
    mapData.visited_countries.forEach(function(location){
      $('#svgContainer').find('path.land').each(function(country){
        var title = $(this).attr('title');
        if(title == location){
          $(this).css({ fill: "#4286f4" });
        }
      });
    });
    populateVistedCountriesTable();
  });
}
