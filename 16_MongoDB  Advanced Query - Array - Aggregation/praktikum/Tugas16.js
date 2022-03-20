//1.Gabungkan (menampilkan) data buku dari author id 1 dan author id 2
db.Books.aggregate([{
    $match:{
    "$or":[{authorID:1},{authorID:2}]
    }
}])


//2.Tampilkan daftar buku dan harga author id 1.
db.Books.aggregate([
    {
    $match:{authorID:1}
    },
    {
     $project:{
      _id:1,
      title:1,
      price:1
      }
    }
])


//3.Tampilan total jumlah halaman buku author id 2.
db.Books.aggregate([
    {
    $match:{authorID:2}
    },
    {
    $group:{
    _id:2,
    totalPage:{$sum:"$_id"}
        }
    }
]);


//4. Tampilkan semua field books and authors terkait.

db.Books.aggregate([
    {
    $lookup:{
       from:'authors',
       localField:"authorID" ,
      foreignField:"_id" ,
     as : 'authors_Name' }
    },
    {
    $project:{
        _id:1,
        title:1,
        "firstName":"$authors_Name.firstName",
        "lastName": "$authors_Name.lastName",
        publisherID:1,
        price:1,
        stats:1,
        publishedAt:1,
        category:1
        }
    }
])



//5. Tampilkan semua field books, authors, dan publishers terkait.

db.Books.aggregate([
    {
    $lookup:{
       from:'authors',
       localField:"authorID" ,
      foreignField:"_id" ,
     as : 'authors_Name' }
    },
    {
    $lookup:{
    from:'publishers',
    localField:"publisherID" ,
      foreignField:"_id" ,
      as : 'publisher_Name' }
    },
    {
    $project:{
        _id:1,
        title:1,
        authorID:1,
        publisherID:1,
        price:1,
        stats:1,
        publishedAt:1,
        category:1,
        "author": "$authors_Name",
        "publisher" : "$publisher_Name"
        }
    }
])


//No 6. Tampilkan summary data authors, books, dan publishers sesuai dengan Output.

db.Books.aggregate([
    {
    $lookup:{
       from:'authors',
       localField:"authorID" ,
      foreignField:"_id" ,
     as : 'authors_Name' }
    },
    {
    $lookup:{
    from:'publishers',
    localField:"publisherID" ,
      foreignField:"_id" ,
      as : 'publisher_Name' }
    },
    
    {
    $group:{
        _id:"$authors_Name.firstName",
        //{$concat:["$authors_Name.firstName"," ", "$authors_Name.lastName"]},
        "numberOfBooks" : {$count:{}},
        "listOfBook" : {$push:{title : "$title",publisher:"$publisher_Name.publisherName"}}

        }
    }
])



//7.Digital_outlet ingin memberikan diskon untuk setiap buku, diskon di tentukan melihat harga buku tersebut dengan pembagian seperti ini.
//Price < 60.000  1 %
//60.000 < Price < 90.000 2 %
//90.000 < Price  3 %

db.Books.aggregate([
    {
    $project:{
        _id:1,
        title:1,
        price:1,
        "Promo" :{
            $cond:{if:{$lt:["$price",60000]},
            then :"1%", else :{
            $cond:{if:{$lt:["$price",90000]},then:"2%",else:"3%"}}
            }
        }
    }
    }  
])



//No 8 Tampilkan semua nama buku, harga, nama author dan nama publisher, urutkan dari harga termahal ke termurah

db.Books.aggregate([
    {
    $lookup:{
       from:'authors',
       localField:"authorID" ,
      foreignField:"_id" ,
     as : 'authors_Name' }
    },
    {
    $lookup:{
    from:'publishers',
    localField:"publisherID" ,
      foreignField:"_id" ,
      as : 'publisher_Name' }
    },
    {
    $project:{
        title:1,
        price:1,
        "namaAuthor":"$authors_Name.firstName",
        //{$concat:["$authors_Name.firstName"," ", "$authors_Name.lastName"]},
        "namaPublisher" : "$publisher_Name.publisherName"
        }
     },
     {
     $sort:{price:1}
     }
])



//9. Tampilkan data nama buku harga dan publisher, kemudian tampilkan hanya data ke 3 dan ke 4.

db.Books.aggregate([
    {
    $lookup:{
    from:'publishers',
    localField:"publisherID" ,
      foreignField:"_id" ,
      as : 'publisher_Name' }
    },
    {
    $project:{
        _id:1,
        title:1,
        price:1,
        "namaPublisher" : "$publisher_Name.publisherName"
        }
     },
     {
     $skip:3
     }
])



