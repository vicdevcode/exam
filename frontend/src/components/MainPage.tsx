import { FC, useState, useEffect } from "react";
import { AddCard } from "./AddCard";
import { Dialog, DialogTrigger } from "./ui/dialog";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "./ui/card";
import { Button } from "./ui/button";
import { BlockMath, InlineMath } from "react-katex";
import { EditCard } from "./EditCard";

interface Category {
  label: string;
  value: number;
}

interface SubCategory {
  label: string;
  value: number;
  category_id: number;
}

interface Cards {
  id: number;
  question: string;
  answer: string;
  sub_category_id: number;
}

const MainPage: FC = () => {
  const [categories, setCategories] = useState<Category[]>([]);
  const [subCategories, setSubCategories] = useState<SubCategory[]>([]);
  const [cards, setCards] = useState<Cards[]>([]);
  const [curCard, setCurCard] = useState<number>(1);
  const [revealCard, setRevealCard] = useState<boolean>(false);

  const getCategories = async () => {
    const res = await fetch("/api/category/all", {
      method: "GET",
    });

    const json = await res.json();

    const c = [];
    const s = [];
    const q = [];

    for (let i = 0; i < json["categories"].length; i++) {
      c.push({
        value: json["categories"][i]["id"],
        label: json["categories"][i]["name"],
      });
      for (let j = 0; j < json["categories"][i]["sub_categories"].length; j++) {
        s.push({
          value: json["categories"][i]["sub_categories"][j]["id"],
          label: json["categories"][i]["sub_categories"][j]["name"],
          category_id:
            json["categories"][i]["sub_categories"][j]["category_id"],
        });
        for (
          let k = 0;
          k < json["categories"][i]["sub_categories"][j]["cards"].length;
          k++
        ) {
          q.push({
            id: json["categories"][i]["sub_categories"][j]["cards"][k]["id"],
            question:
              json["categories"][i]["sub_categories"][j]["cards"][k][
              "question"
              ],
            answer:
              json["categories"][i]["sub_categories"][j]["cards"][k]["answer"],
            sub_category_id:
              json["categories"][i]["sub_categories"][j]["cards"][k][
              "sub_category_id"
              ],
          });
        }
      }
    }
    setCategories(c);
    setSubCategories(s);
    setCards(q);
  };

  const processText = (text: string) => {
    const parts = text.split(/(\$\$.*?\$\$|\$.*?\$|\n)/g).filter(Boolean);
    return parts.map((part, index) => {
      if (part === "\n") {
        return <br key={index} />;
      } else if (part.startsWith("$$") && part.endsWith("$$")) {
        const math = part.slice(2, -2).replace(/\\\\/g, "\\");
        return <BlockMath key={index} math={math} />;
      } else if (part.startsWith("$") && part.endsWith("$")) {
        const math = part.slice(1, -1).replace(/\\\\/g, "\\");
        return <InlineMath key={index} math={math} />;
      } else {
        return part;
      }
    });
    //const parts = text.split(/(\$\$.*?\$\$|\$.*?\$)/g).filter(Boolean);
    //return parts.map((part, index) => {
    //  if (part.startsWith("$$") && part.endsWith("$$")) {
    //    const math = part.slice(2, -2).replace(/\\\\/g, "\\");
    //    return <BlockMath key={index} math={math} />;
    //  } else if (part.startsWith("$") && part.endsWith("$")) {
    //    const math = part.slice(1, -1).replace(/\\\\/g, "\\");
    //    return <InlineMath key={index} math={math} />;
    //  } else {
    //    return part;
    //  }
    //});
  };

  const next = () => {
    setCurCard((c) => (c == cards.length ? 1 : c + 1));
    setRevealCard(false);
  };

  useEffect(() => {
    getCategories();
  }, []);

  return (
    <div className="h-screen">
      <header>
        <Dialog>
          <DialogTrigger>Добавить карточку</DialogTrigger>
          <AddCard subCategories={subCategories} />
        </Dialog>
      </header>
      <div className="flex justify-center items-center pb-[200px] h-screen m-auto w-2/3">
        <div className="w-full">
          {cards.length > 0 && (
            <Card>
              <CardHeader>
                <div className="flex justify-between text-sm">
                  <span>
                    {
                      categories.find(
                        (c) =>
                          c.value ==
                          subCategories.find(
                            (e) =>
                              e.value == cards[curCard - 1]["sub_category_id"],
                          )?.category_id,
                      )?.label
                    }
                  </span>
                  <span>
                    {
                      subCategories.find(
                        (e) => e.value == cards[curCard - 1]["sub_category_id"],
                      )?.label
                    }
                  </span>
                </div>
                <CardTitle className="text-center">
                  {curCard}/{cards.length}
                </CardTitle>
              </CardHeader>
              <CardContent className="py-10 flex justify-center items-center">
                {revealCard ? (
                  <div className="flex flex-col justify-center items-center gap-3">
                    <span className="text-lg font-medium">
                      {cards[curCard - 1].question}
                    </span>
                    <span className="text-xl font-medium">
                      {processText(cards[curCard - 1].answer)}
                    </span>
                  </div>
                ) : (
                  <span className="text-3xl font-medium">
                    {cards[curCard - 1].question}
                  </span>
                )}
              </CardContent>
              <CardFooter className="pt-6 flex justify-between">
                {revealCard ? (
                  <>
                    <Dialog>
                      <DialogTrigger>Отредактировать</DialogTrigger>
                      <EditCard
                        id={cards[curCard - 1].id}
                        sub_category_id={cards[curCard - 1].sub_category_id}
                        question={cards[curCard - 1].question}
                        answer={cards[curCard - 1].answer}
                        subCategories={subCategories}
                      />
                    </Dialog>
                    <Button onClick={next}>Понятно</Button>
                  </>
                ) : (
                  <>
                    <Button
                      onClick={() => setRevealCard(true)}
                      variant="outline"
                    >
                      Не знаю ответа
                    </Button>
                    <Button onClick={next}>Знаю ответ</Button>
                  </>
                )}
              </CardFooter>
            </Card>
          )}
        </div>
      </div>
    </div>
  );
};

export default MainPage;
