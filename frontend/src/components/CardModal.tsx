import { FC } from "react";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogTrigger,
} from "./ui/dialog";
import { BlockMath, InlineMath } from "react-katex";
import { EditCard } from "./EditCard";
import { Button } from "./ui/button";

interface CardModalProps {
  id: number;
  question: string;
  answer: string;
  sub_category_id: number;
  subCategories: {
    label: string;
    value: number;
    category_id: number;
  }[];
}

export const CardModal: FC<CardModalProps> = (props) => {
  const { id, question, answer, sub_category_id, subCategories } = props;

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
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="link" size="sm">
          {question}
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-[1000px] px-4">
        <div className="flex flex-col justify-center items-center gap-3">
          <span className="text-lg font-medium">{processText(question)}</span>
          <span className="text-xl font-medium">{processText(answer)}</span>
        </div>
        <DialogFooter>
          <Dialog>
            <DialogTrigger>Отредактировать</DialogTrigger>
            <EditCard
              id={id}
              sub_category_id={sub_category_id}
              question={question}
              answer={answer}
              subCategories={subCategories}
            />
          </Dialog>
          <DialogClose asChild>
            <Button>Понятно</Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};
